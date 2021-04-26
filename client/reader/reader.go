package reader

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/arxdsilva/jpar/client/domains"
	"github.com/arxdsilva/jpar/client/infrastructure/grpc_client"
	"github.com/kpango/glg"
)

func StreamFile(semaphore chan domains.Port) (err error) {
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
		// read ID "AEAJM"
		t, err := decoder.Token()
		glg.Info("[streamFile] token, err: ", t, err)
		c := &domains.Port{}
		err = decoder.Decode(c)
		if err != nil {
			glg.Info("[streamFile] decode err: ", err.Error())
		}
		c.ID = t.(string)
		glg.Info("[streamFile] CITY: ", c.City, c.ID)
		semaphore <- *c
	}
	token, err = decoder.Token()
	if err != nil {
		return
	}
	glg.Info("[streamFile] finished parsing file, last token: ", token)
	return
}

func SendInfo(semaphore chan domains.Port) {
	for {
		port, open := <-semaphore
		if !open {
			return
		}
		fmt.Println("send port: ", port.Code, port.Province)
		go func() {
			grpc_client.SendPortToServer(port)
		}()
	}
}
