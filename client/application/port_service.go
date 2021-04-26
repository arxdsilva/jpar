package application

import (
	"github.com/arxdsilva/jpar/client/domains"
	"github.com/arxdsilva/jpar/client/infrastructure/grpc_client"
)

type portService struct{}

func NewPortService() domains.PortService {
	return &portService{}
}

func (p *portService) GetPorts() (ps []domains.Port, err error) {
	return grpc_client.GetAllPorts()
}
