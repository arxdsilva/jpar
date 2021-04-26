package grpc_client

import (
	"context"
	"log"
	"time"

	"github.com/arxdsilva/jpar/client/domains"
	pb "github.com/arxdsilva/jpar/port"
	"github.com/kpango/glg"
	"google.golang.org/grpc"
)

// this should be refactored into a struct that doesnt open 1 conn to each request
func SendPortToServer(cp domains.Port) {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
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
