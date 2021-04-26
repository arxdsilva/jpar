package reader

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/arxdsilva/jpar/server/port"
	"github.com/kpango/glg"
	"google.golang.org/grpc"
)

type CityPort struct {
	PortID      string        `json:"port_id"`
	Name        string        `json:"name"`
	Coordinates []float64     `json:"coordinates"`
	City        string        `json:"city"`
	Province    string        `json:"province"`
	Country     string        `json:"country"`
	Alias       []string 	  `json:"alias"`
	Regions     []string 	  `json:"regions"`
	Timezone    string        `json:"timezone"`
	Unlocs      []string      `json:"unlocs"`
	Code        string        `json:"code"`
}

func StreamFile(semaphore chan CityPort) (err error) {
	defer close(semaphore)
	file, err := os.Open("ports.json")
	if err != nil {
		glg.Error("Failed to load file: %v", err)
		return
	}
	decoder := json.NewDecoder(file)
	// read { opener
	token, err := decoder.Token()
	if err != nil {
		return
	}
	glg.Info("[streamFile] started parsing file, first token: ", token)
	for decoder.More() {
		// read PortID "AEAJM"
		t, err := decoder.Token()
		glg.Info("[streamFile] token1, err: ", t, err)
		c := &CityPort{}
		err = decoder.Decode(c)
		if err != nil {
			glg.Info("[streamFile] decode err: ", err.Error())
		}
		c.PortID = t.(string)
		glg.Info("[streamFile] CITY: ", c.City, c.PortID)
		semaphore <- *c
	}
	token, err = decoder.Token()
	if err != nil {
		return
	}
	glg.Info("[streamFile] finished parsing file, last token: ", token)
	return
}

func SendInfo(semaphore chan CityPort) {
	for {
		port, open := <-semaphore
		if !open {
			return
		}
		fmt.Println("send port: ", port.Code, port.Province)
		go func ()  {
			sendPortToServer(port)
		}
	}
}

func sendPortToServer(cp *CityPort) {
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
		PortId:      cp.PortID,
	}
	resp, err := c.UpsertPort(ctx, pds)
	if err != nil {
		glg.Error("[sendPortToServer] err ", err.Error())
		return
	}
	if resp.Error != "" {
		return
	}
	glg.Info("[sendPortToServer] ok ", cp.PortID)
	return
}
