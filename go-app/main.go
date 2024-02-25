package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.000",
	})
}

func Info(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var output = " | " + getEnv() + " | " + currentTime.Format("2006.01.02 15:04:05.000000") + " | " + hostname + " | "
	io.WriteString(w, output)
	log.Info(output)
}

func getEnv() string {
	value := os.Getenv("APP_VERSION")
	if len(value) == 0 {
		return "0.0.1"
	}
	return value
}

func getFrontpage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Congratulations! Version %s of your application is running on Kubernetes.", getEnv())
	log.Info("Congratulations! Version " + getEnv() + " of your application is running on Kubernetes.")
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func getVersion(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", getEnv())
	log.Info(getEnv())
}

func main() {
	http.HandleFunc("/go-demo", getFrontpage)
	http.HandleFunc("/go-demo/health", health)
	http.HandleFunc("/go-demo/info", Info)
	http.HandleFunc("/go-demo/version", getVersion)
	log.Info("********** Go Front Started **********")
	http.ListenAndServe(":8888", nil)
}
