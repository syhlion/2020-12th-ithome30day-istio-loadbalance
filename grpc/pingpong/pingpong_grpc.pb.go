// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pingpong

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PingPongServiceClient is the client API for PingPongService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PingPongServiceClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PongResponse, error)
}

type pingPongServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPingPongServiceClient(cc grpc.ClientConnInterface) PingPongServiceClient {
	return &pingPongServiceClient{cc}
}

func (c *pingPongServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PongResponse, error) {
	out := new(PongResponse)
	err := c.cc.Invoke(ctx, "/pingpong.PingPongService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PingPongServiceServer is the server API for PingPongService service.
// All implementations must embed UnimplementedPingPongServiceServer
// for forward compatibility
type PingPongServiceServer interface {
	Ping(context.Context, *PingRequest) (*PongResponse, error)
	mustEmbedUnimplementedPingPongServiceServer()
}

// UnimplementedPingPongServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPingPongServiceServer struct {
}

func (*UnimplementedPingPongServiceServer) Ping(context.Context, *PingRequest) (*PongResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedPingPongServiceServer) mustEmbedUnimplementedPingPongServiceServer() {}

func RegisterPingPongServiceServer(s *grpc.Server, srv PingPongServiceServer) {
	s.RegisterService(&_PingPongService_serviceDesc, srv)
}

func _PingPongService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingPongServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pingpong.PingPongService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingPongServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PingPongService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pingpong.PingPongService",
	HandlerType: (*PingPongServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _PingPongService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pingpong.proto",
}
