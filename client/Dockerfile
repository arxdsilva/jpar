FROM golang:latest AS build
WORKDIR /src
ENV GO111MODULE=on
COPY . .
RUN go mod download
RUN go build -o api ./cmd/api/main.go
EXPOSE 8888
ENTRYPOINT ./api