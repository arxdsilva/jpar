package main

import (
	"net"

	"github.com/arxdsilva/jpar/backend/infrastructure/config"
	"github.com/kpango/glg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config.Load()
	listen, err := net.Listen("tcp", ":8083")
	if err != nil {
		glg.Fatal(err.Error())
	}
	server := grpc.NewServer()
	reflection.Register(server)
	glg.Info("starting server at :8083")
	err = server.Serve(listen)
	if err != nil {
		glg.Fatal(err.Error())
	}
}
