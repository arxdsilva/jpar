package main

import (
	"net"

	"github.com/arxdsilva/jpar/backend/infrastructure/config"
	infraGRPC "github.com/arxdsilva/jpar/backend/infrastructure/grpc"
	"github.com/arxdsilva/jpar/backend/infrastructure/repository"
	"github.com/arxdsilva/jpar/backend/interfaces"
	pb "github.com/arxdsilva/jpar/client/port"
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
	repo := repository.NewService()
	portService := interfaces.NewPortService(repo)
	pb.RegisterPortDomainServiceServer(server, infraGRPC.NewPortServer(portService))
	reflection.Register(server)
	glg.Info("starting server at :8083")
	err = server.Serve(listen)
	if err != nil {
		glg.Fatal(err.Error())
	}
}
