// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package port

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PortDomainServiceClient is the client API for PortDomainService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PortDomainServiceClient interface {
	UpsertPort(ctx context.Context, in *Port, opts ...grpc.CallOption) (*PortResponse, error)
	ListPorts(ctx context.Context, in *List, opts ...grpc.CallOption) (PortDomainService_ListPortsClient, error)
}

type portDomainServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPortDomainServiceClient(cc grpc.ClientConnInterface) PortDomainServiceClient {
	return &portDomainServiceClient{cc}
}

func (c *portDomainServiceClient) UpsertPort(ctx context.Context, in *Port, opts ...grpc.CallOption) (*PortResponse, error) {
	out := new(PortResponse)
	err := c.cc.Invoke(ctx, "/routeguide.PortDomainService/UpsertPort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portDomainServiceClient) ListPorts(ctx context.Context, in *List, opts ...grpc.CallOption) (PortDomainService_ListPortsClient, error) {
	stream, err := c.cc.NewStream(ctx, &PortDomainService_ServiceDesc.Streams[0], "/routeguide.PortDomainService/ListPorts", opts...)
	if err != nil {
		return nil, err
	}
	x := &portDomainServiceListPortsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PortDomainService_ListPortsClient interface {
	Recv() (*Port, error)
	grpc.ClientStream
}

type portDomainServiceListPortsClient struct {
	grpc.ClientStream
}

func (x *portDomainServiceListPortsClient) Recv() (*Port, error) {
	m := new(Port)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PortDomainServiceServer is the server API for PortDomainService service.
// All implementations must embed UnimplementedPortDomainServiceServer
// for forward compatibility
type PortDomainServiceServer interface {
	UpsertPort(context.Context, *Port) (*PortResponse, error)
	ListPorts(*List, PortDomainService_ListPortsServer) error
	mustEmbedUnimplementedPortDomainServiceServer()
}

// UnimplementedPortDomainServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPortDomainServiceServer struct {
}

func (UnimplementedPortDomainServiceServer) UpsertPort(context.Context, *Port) (*PortResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertPort not implemented")
}
func (UnimplementedPortDomainServiceServer) ListPorts(*List, PortDomainService_ListPortsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListPorts not implemented")
}
func (UnimplementedPortDomainServiceServer) mustEmbedUnimplementedPortDomainServiceServer() {}

// UnsafePortDomainServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PortDomainServiceServer will
// result in compilation errors.
type UnsafePortDomainServiceServer interface {
	mustEmbedUnimplementedPortDomainServiceServer()
}

func RegisterPortDomainServiceServer(s grpc.ServiceRegistrar, srv PortDomainServiceServer) {
	s.RegisterService(&PortDomainService_ServiceDesc, srv)
}

func _PortDomainService_UpsertPort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Port)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortDomainServiceServer).UpsertPort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/routeguide.PortDomainService/UpsertPort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortDomainServiceServer).UpsertPort(ctx, req.(*Port))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortDomainService_ListPorts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(List)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PortDomainServiceServer).ListPorts(m, &portDomainServiceListPortsServer{stream})
}

type PortDomainService_ListPortsServer interface {
	Send(*Port) error
	grpc.ServerStream
}

type portDomainServiceListPortsServer struct {
	grpc.ServerStream
}

func (x *portDomainServiceListPortsServer) Send(m *Port) error {
	return x.ServerStream.SendMsg(m)
}

// PortDomainService_ServiceDesc is the grpc.ServiceDesc for PortDomainService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PortDomainService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "routeguide.PortDomainService",
	HandlerType: (*PortDomainServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpsertPort",
			Handler:    _PortDomainService_UpsertPort_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListPorts",
			Handler:       _PortDomainService_ListPorts_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "port/port.proto",
}
