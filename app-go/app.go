package main

import (
	"fmt"
	"net/http"
	"os"
)

var version string = "N/A"

func getEnv() {
	version = os.Getenv("APP_VERSION")
}

func getFrontpage(w http.ResponseWriter, r *http.Request) {
	getEnv()
	fmt.Printf("Congratulations! Version %s of your application is running on Kubernetes.\n", version)
	fmt.Fprintf(w, "Congratulations! Version %s of your application is running on Kubernetes.", version)
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func getVersion(w http.ResponseWriter, r *http.Request) {
	getEnv()
	fmt.Printf("%s\n", version)
	fmt.Fprintf(w, "%s\n", version)
}

func main() {
	http.HandleFunc("/", getFrontpage)
	http.HandleFunc("/health", health)
	http.HandleFunc("/version", getVersion)
	http.ListenAndServe(":8080", nil)
}
