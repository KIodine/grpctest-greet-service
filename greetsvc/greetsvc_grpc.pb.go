// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package greetsvc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	greet "grpcb/greet"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GreetingServiceClient is the client API for GreetingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreetingServiceClient interface {
	GreetOnce(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*greet.Greet, error)
	GreetAtOnce(ctx context.Context, opts ...grpc.CallOption) (GreetingService_GreetAtOnceClient, error)
	GreetMany(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (GreetingService_GreetManyClient, error)
	GreetStream(ctx context.Context, opts ...grpc.CallOption) (GreetingService_GreetStreamClient, error)
}

type greetingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetingServiceClient(cc grpc.ClientConnInterface) GreetingServiceClient {
	return &greetingServiceClient{cc}
}

func (c *greetingServiceClient) GreetOnce(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*greet.Greet, error) {
	out := new(greet.Greet)
	err := c.cc.Invoke(ctx, "/greetsvc.GreetingService/GreetOnce", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greetingServiceClient) GreetAtOnce(ctx context.Context, opts ...grpc.CallOption) (GreetingService_GreetAtOnceClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetingService_ServiceDesc.Streams[0], "/greetsvc.GreetingService/GreetAtOnce", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetingServiceGreetAtOnceClient{stream}
	return x, nil
}

type GreetingService_GreetAtOnceClient interface {
	Send(*GreetRequest) error
	CloseAndRecv() (*GreetGathered, error)
	grpc.ClientStream
}

type greetingServiceGreetAtOnceClient struct {
	grpc.ClientStream
}

func (x *greetingServiceGreetAtOnceClient) Send(m *GreetRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetingServiceGreetAtOnceClient) CloseAndRecv() (*GreetGathered, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(GreetGathered)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetingServiceClient) GreetMany(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (GreetingService_GreetManyClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetingService_ServiceDesc.Streams[1], "/greetsvc.GreetingService/GreetMany", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetingServiceGreetManyClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GreetingService_GreetManyClient interface {
	Recv() (*greet.Greet, error)
	grpc.ClientStream
}

type greetingServiceGreetManyClient struct {
	grpc.ClientStream
}

func (x *greetingServiceGreetManyClient) Recv() (*greet.Greet, error) {
	m := new(greet.Greet)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetingServiceClient) GreetStream(ctx context.Context, opts ...grpc.CallOption) (GreetingService_GreetStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetingService_ServiceDesc.Streams[2], "/greetsvc.GreetingService/GreetStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetingServiceGreetStreamClient{stream}
	return x, nil
}

type GreetingService_GreetStreamClient interface {
	Send(*GreetRequest) error
	Recv() (*greet.Greet, error)
	grpc.ClientStream
}

type greetingServiceGreetStreamClient struct {
	grpc.ClientStream
}

func (x *greetingServiceGreetStreamClient) Send(m *GreetRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetingServiceGreetStreamClient) Recv() (*greet.Greet, error) {
	m := new(greet.Greet)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetingServiceServer is the server API for GreetingService service.
// All implementations must embed UnimplementedGreetingServiceServer
// for forward compatibility
type GreetingServiceServer interface {
	GreetOnce(context.Context, *GreetRequest) (*greet.Greet, error)
	GreetAtOnce(GreetingService_GreetAtOnceServer) error
	GreetMany(*GreetRequest, GreetingService_GreetManyServer) error
	GreetStream(GreetingService_GreetStreamServer) error
	mustEmbedUnimplementedGreetingServiceServer()
}

// UnimplementedGreetingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGreetingServiceServer struct {
}

func (UnimplementedGreetingServiceServer) GreetOnce(context.Context, *GreetRequest) (*greet.Greet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GreetOnce not implemented")
}
func (UnimplementedGreetingServiceServer) GreetAtOnce(GreetingService_GreetAtOnceServer) error {
	return status.Errorf(codes.Unimplemented, "method GreetAtOnce not implemented")
}
func (UnimplementedGreetingServiceServer) GreetMany(*GreetRequest, GreetingService_GreetManyServer) error {
	return status.Errorf(codes.Unimplemented, "method GreetMany not implemented")
}
func (UnimplementedGreetingServiceServer) GreetStream(GreetingService_GreetStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GreetStream not implemented")
}
func (UnimplementedGreetingServiceServer) mustEmbedUnimplementedGreetingServiceServer() {}

// UnsafeGreetingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreetingServiceServer will
// result in compilation errors.
type UnsafeGreetingServiceServer interface {
	mustEmbedUnimplementedGreetingServiceServer()
}

func RegisterGreetingServiceServer(s grpc.ServiceRegistrar, srv GreetingServiceServer) {
	s.RegisterService(&GreetingService_ServiceDesc, srv)
}

func _GreetingService_GreetOnce_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetingServiceServer).GreetOnce(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greetsvc.GreetingService/GreetOnce",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetingServiceServer).GreetOnce(ctx, req.(*GreetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreetingService_GreetAtOnce_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetingServiceServer).GreetAtOnce(&greetingServiceGreetAtOnceServer{stream})
}

type GreetingService_GreetAtOnceServer interface {
	SendAndClose(*GreetGathered) error
	Recv() (*GreetRequest, error)
	grpc.ServerStream
}

type greetingServiceGreetAtOnceServer struct {
	grpc.ServerStream
}

func (x *greetingServiceGreetAtOnceServer) SendAndClose(m *GreetGathered) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetingServiceGreetAtOnceServer) Recv() (*GreetRequest, error) {
	m := new(GreetRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GreetingService_GreetMany_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GreetRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreetingServiceServer).GreetMany(m, &greetingServiceGreetManyServer{stream})
}

type GreetingService_GreetManyServer interface {
	Send(*greet.Greet) error
	grpc.ServerStream
}

type greetingServiceGreetManyServer struct {
	grpc.ServerStream
}

func (x *greetingServiceGreetManyServer) Send(m *greet.Greet) error {
	return x.ServerStream.SendMsg(m)
}

func _GreetingService_GreetStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetingServiceServer).GreetStream(&greetingServiceGreetStreamServer{stream})
}

type GreetingService_GreetStreamServer interface {
	Send(*greet.Greet) error
	Recv() (*GreetRequest, error)
	grpc.ServerStream
}

type greetingServiceGreetStreamServer struct {
	grpc.ServerStream
}

func (x *greetingServiceGreetStreamServer) Send(m *greet.Greet) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetingServiceGreetStreamServer) Recv() (*GreetRequest, error) {
	m := new(GreetRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetingService_ServiceDesc is the grpc.ServiceDesc for GreetingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GreetingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "greetsvc.GreetingService",
	HandlerType: (*GreetingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GreetOnce",
			Handler:    _GreetingService_GreetOnce_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GreetAtOnce",
			Handler:       _GreetingService_GreetAtOnce_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GreetMany",
			Handler:       _GreetingService_GreetMany_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GreetStream",
			Handler:       _GreetingService_GreetStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "greetsvc/greetsvc.proto",
}
