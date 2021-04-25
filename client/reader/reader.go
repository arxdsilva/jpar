package reader

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

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
	Alias       []interface{} `json:"alias"`
	Regions     []interface{} `json:"regions"`
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
		city, open := <-semaphore
		if !open {
			return
		}
		fmt.Println("send city: ", city.Code, city.Province)
	}
}

func sendPortToServer(cp *CityPort) {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// c := pb.NewGreeterClient(conn)
	// // Contact the server and print out its response.
	// name := defaultName
	// if len(os.Args) > 1 {
	// 	name = os.Args[1]
	// }
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()
	// r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	// if err != nil {
	// 	log.Fatalf("could not greet: %v", err)
	// }
	// log.Printf("Greeting: %s", r.GetMessage())
	return
}
