package grpc

import (
	"context"

	"github.com/arxdsilva/jpar/backend/domain"
	pb "github.com/arxdsilva/jpar/client/port"
	"github.com/kpango/glg"
	"google.golang.org/grpc"
)

type portServer struct {
	svc domain.PortService
}

func NewPortServer(svc domain.PortService) *portServer {
	return &portServer{svc: svc}
}

func (s *portServer) UpsertPort(ctx context.Context, in *pb.Port, opts ...grpc.CallOption) (pr *pb.PortResponse, err error) {
	glg.Info("[UpsertPort] start")
	port := domain.Port{
		ID:          in.ID,
		Name:        in.Name,
		Coordinates: in.Coordinates,
		City:        in.City,
		Province:    in.Province,
		Country:     in.Country,
		Alias:       in.Alias,
		Regions:     in.Regions,
		Timezone:    in.Timezone,
		Unlocs:      in.Unlocs,
		Code:        in.Code,
	}
	err = s.svc.UpsertPort(port)
	if err != nil {
		pr.Error = err.Error()
		return
	}
	glg.Info("[UpsertPort] finish")
	return
}

func (s *portServer) ListPorts(ctx context.Context, in *pb.List, opts ...grpc.CallOption) (lpc pb.PortDomainService_ListPortsClient, err error) {
	glg.Info("[ListPorts] start")
	ports, err := s.svc.ListPorts()
	if err != nil {
		lpc.Error = err.Error()
		return
	}
	server := pb.PortDomainService_ListPortsServer{}
	for _, p := range ports {
		server.Send(&p)
	}
	glg.Info("[ListPorts] finish")
	return
}
