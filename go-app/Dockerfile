# stage build
FROM golang:latest as build

WORKDIR /build

COPY go.mod ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /build/app main.go

# stage imagem final
FROM alpine:latest

WORKDIR /apps

COPY --from=build /build/app ./

EXPOSE 8080

CMD [ "./app" ]