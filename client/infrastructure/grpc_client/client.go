package grpc_client

import (
	"context"
	"io"
	"log"
	"os"
	"time"

	"github.com/arxdsilva/jpar/client/domains"
	pb "github.com/arxdsilva/jpar/client/port"
	"github.com/kpango/glg"
	"google.golang.org/grpc"
)

// this should be refactored into a struct that doesnt open 1 conn to each request
func SendPortToServer(cp domains.Port) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(os.Getenv("BACKEND_URI"), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPortDomainServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	pds := pb.Port{
		Name:        cp.Name,
		City:        cp.City,
		Country:     cp.Country,
		Alias:       cp.Alias,
		Regions:     cp.Regions,
		Coordinates: cp.Coordinates,
		Province:    cp.Province,
		Timezone:    cp.Timezone,
		Unlocs:      cp.Unlocs,
		Code:        cp.Code,
		Id:          cp.ID,
	}
	resp, err := c.UpsertPort(ctx, &pds)
	if err != nil {
		glg.Error("[sendPortToServer] err ", err.Error())
		return
	}
	if resp.Error != "" {
		return
	}
	glg.Info("[sendPortToServer] ok ", cp.ID)
	return
}

func GetAllPorts() (ps []domains.Port, err error) {
	ps = []domains.Port{}
	conn, err := grpc.Dial(os.Getenv("BACKEND_URI"), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPortDomainServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	list := &pb.List{}
	listener, err := c.ListPorts(ctx, list)
	timeout := time.NewTicker(30 * time.Second)
	for {
		select {
		case <-timeout.C:
			err = listener.CloseSend()
			if err != nil {
				glg.Error("[GetAllPorts] error closing listener: ", err.Error())
				return ps, nil
			}
			return
		default:
			resp, errL := listener.Recv()
			if errL != nil && err != io.EOF {
				glg.Error("[GetAllPorts] error closing listener: ", errL.Error())
				return ps, errL
			}
			if err == io.EOF {
				glg.Info("[GetAllPorts] stream finished: ")
				return
			}
			port := domains.Port{
				ID:          resp.Id,
				Name:        resp.Name,
				Coordinates: resp.Coordinates,
				City:        resp.City,
				Province:    resp.Province,
				Country:     resp.Country,
				Alias:       resp.Alias,
				Regions:     resp.Regions,
				Timezone:    resp.Timezone,
				Unlocs:      resp.Unlocs,
				Code:        resp.Code,
			}
			ps = append(ps, port)
		}
	}
}
