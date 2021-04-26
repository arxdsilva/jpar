package interfaces

import (
	"encoding/json"
	"os"

	"github.com/arxdsilva/jpar/client/domains"
	"github.com/arxdsilva/jpar/client/infrastructure/config"
	"github.com/arxdsilva/jpar/client/infrastructure/grpc_client"
	"github.com/kpango/glg"
)

func streamFile(c config.Config) (err error) {
	defer close(c.Semaphore)
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
		port := &domains.Port{}
		err = decoder.Decode(port)
		if err != nil {
			glg.Info("[streamFile] decode err: ", err.Error())
		}
		port.ID = t.(string)
		glg.Info("[streamFile] CITY: ", port.City, port.ID)
		c.Semaphore <- *port
	}
	token, err = decoder.Token()
	if err != nil {
		return
	}
	glg.Info("[streamFile] finished parsing file, last token: ", token)
	return
}

func sendPortData(c config.Config) {
	for {
		port, open := <-c.Semaphore
		if !open {
			return
		}
		glg.Info("[sendPortData] ", port.ID)
		go func() { grpc_client.SendPortToServer(port) }()
	}
}
