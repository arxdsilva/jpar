package grpc

import (
	"context"

	"github.com/arxdsilva/jpar/backend/domain"
	pb "github.com/arxdsilva/jpar/client/port"
	"google.golang.org/grpc"
)

type portServer struct {
	svc domain.PortService
}

func NewPortServer(svc domain.PortService) *portServer {
	return &portServer{}
}

func (s *portServer) UpsertPort(ctx context.Context, in *pb.Port, opts ...grpc.CallOption) (pr *pb.PortResponse, err error) {

	return
}

func (s *portServer) ListPorts(ctx context.Context, in *pb.List, opts ...grpc.CallOption) (lpc pb.PortDomainService_ListPortsClient, err error) {

	return
}
