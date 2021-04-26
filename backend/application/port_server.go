package application

import (
	"context"

	"github.com/arxdsilva/jpar/backend/domain"
	pb "github.com/arxdsilva/jpar/client/port"
	"github.com/kpango/glg"
)

type portServer struct {
	svc domain.PortService
	pb.UnimplementedPortDomainServiceServer
}

func NewPortServer(svc domain.PortService) *portServer {
	return &portServer{svc: svc}
}

func (s *portServer) UpsertPort(ctx context.Context, in *pb.Port) (pr *pb.PortResponse, err error) {
	glg.Info("[UpsertPort] start")
	port := domain.Port{
		ID:          in.Id,
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

func (s *portServer) ListPorts(in *pb.List, stream pb.PortDomainService_ListPortsServer) (err error) {
	glg.Info("[ListPorts] start")
	ports, err := s.svc.ListPorts()
	if err != nil {
		return
	}
	for _, p := range ports {
		port := &pb.Port{
			Id:          p.ID,
			Name:        p.Name,
			Coordinates: p.Coordinates,
			City:        p.City,
			Province:    p.Province,
			Country:     p.Country,
			Alias:       p.Alias,
			Regions:     p.Regions,
			Timezone:    p.Timezone,
			Unlocs:      p.Unlocs,
			Code:        p.Code,
		}
		stream.Send(port)
	}
	glg.Info("[ListPorts] finish")
	return
}
