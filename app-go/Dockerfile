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