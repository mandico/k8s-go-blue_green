# k8s | go | blue_green

### Structure Project
``` shell
.
├── README.md
├── app-go             --- Folder Project
│   ├── Dockerfile     --- Dockerfile
│   ├── app.go
│   └── go.mod
└── k8s                --- Folder Manifests Kubernetes
    ├── blue.yaml
    ├── green.yaml
    ├── ingress.yaml
    └── svc-lb.yaml
```

### Build Application
``` shell
go mod download && go mod verify
CGO_ENABLED=0 GOOS=linux go build -o /build/app app.go
```

### Dockerfile
``` Dockerfile
# stage build
FROM golang:1.19 as build

WORKDIR /build

COPY go.mod ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /build/app app.go

# stage imagem final
FROM alpine:3.16

WORKDIR /apps

COPY --from=build /build/app ./

EXPOSE 8080

CMD [ "./app" ]
```

### Build App + Build Image
``` shell
cd app-go/
docker build . -t go-app:1.0.0 -t go-app:latest
```

### Deployment
``` shell
cd k8s/
kubectl apply -f .
```


 while true; do curl http://kubernetes.docker.internal/go-demo/version && sleep 1; done