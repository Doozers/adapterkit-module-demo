// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package demomod

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

// DemomodSvcClient is the client API for DemomodSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DemomodSvcClient interface {
	Sum(ctx context.Context, in *Sum_Request, opts ...grpc.CallOption) (*Sum_Result, error)
	SayHello(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HelloResult, error)
	EchoStream(ctx context.Context, opts ...grpc.CallOption) (DemomodSvc_EchoStreamClient, error)
}

type demomodSvcClient struct {
	cc grpc.ClientConnInterface
}

func NewDemomodSvcClient(cc grpc.ClientConnInterface) DemomodSvcClient {
	return &demomodSvcClient{cc}
}

func (c *demomodSvcClient) Sum(ctx context.Context, in *Sum_Request, opts ...grpc.CallOption) (*Sum_Result, error) {
	out := new(Sum_Result)
	err := c.cc.Invoke(ctx, "/io.moul.adapterkit_module_demo.DemomodSvc/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demomodSvcClient) SayHello(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HelloResult, error) {
	out := new(HelloResult)
	err := c.cc.Invoke(ctx, "/io.moul.adapterkit_module_demo.DemomodSvc/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demomodSvcClient) EchoStream(ctx context.Context, opts ...grpc.CallOption) (DemomodSvc_EchoStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &DemomodSvc_ServiceDesc.Streams[0], "/io.moul.adapterkit_module_demo.DemomodSvc/EchoStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &demomodSvcEchoStreamClient{stream}
	return x, nil
}

type DemomodSvc_EchoStreamClient interface {
	Send(*EchoStream_Input) error
	Recv() (*EchoStream_Output, error)
	grpc.ClientStream
}

type demomodSvcEchoStreamClient struct {
	grpc.ClientStream
}

func (x *demomodSvcEchoStreamClient) Send(m *EchoStream_Input) error {
	return x.ClientStream.SendMsg(m)
}

func (x *demomodSvcEchoStreamClient) Recv() (*EchoStream_Output, error) {
	m := new(EchoStream_Output)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DemomodSvcServer is the server API for DemomodSvc service.
// All implementations must embed UnimplementedDemomodSvcServer
// for forward compatibility
type DemomodSvcServer interface {
	Sum(context.Context, *Sum_Request) (*Sum_Result, error)
	SayHello(context.Context, *Empty) (*HelloResult, error)
	EchoStream(DemomodSvc_EchoStreamServer) error
	mustEmbedUnimplementedDemomodSvcServer()
}

// UnimplementedDemomodSvcServer must be embedded to have forward compatible implementations.
type UnimplementedDemomodSvcServer struct {
}

func (UnimplementedDemomodSvcServer) Sum(context.Context, *Sum_Request) (*Sum_Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (UnimplementedDemomodSvcServer) SayHello(context.Context, *Empty) (*HelloResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedDemomodSvcServer) EchoStream(DemomodSvc_EchoStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method EchoStream not implemented")
}
func (UnimplementedDemomodSvcServer) mustEmbedUnimplementedDemomodSvcServer() {}

// UnsafeDemomodSvcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DemomodSvcServer will
// result in compilation errors.
type UnsafeDemomodSvcServer interface {
	mustEmbedUnimplementedDemomodSvcServer()
}

func RegisterDemomodSvcServer(s grpc.ServiceRegistrar, srv DemomodSvcServer) {
	s.RegisterService(&DemomodSvc_ServiceDesc, srv)
}

func _DemomodSvc_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Sum_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemomodSvcServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.moul.adapterkit_module_demo.DemomodSvc/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemomodSvcServer).Sum(ctx, req.(*Sum_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _DemomodSvc_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemomodSvcServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.moul.adapterkit_module_demo.DemomodSvc/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemomodSvcServer).SayHello(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DemomodSvc_EchoStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DemomodSvcServer).EchoStream(&demomodSvcEchoStreamServer{stream})
}

type DemomodSvc_EchoStreamServer interface {
	Send(*EchoStream_Output) error
	Recv() (*EchoStream_Input, error)
	grpc.ServerStream
}

type demomodSvcEchoStreamServer struct {
	grpc.ServerStream
}

func (x *demomodSvcEchoStreamServer) Send(m *EchoStream_Output) error {
	return x.ServerStream.SendMsg(m)
}

func (x *demomodSvcEchoStreamServer) Recv() (*EchoStream_Input, error) {
	m := new(EchoStream_Input)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DemomodSvc_ServiceDesc is the grpc.ServiceDesc for DemomodSvc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DemomodSvc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "io.moul.adapterkit_module_demo.DemomodSvc",
	HandlerType: (*DemomodSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _DemomodSvc_Sum_Handler,
		},
		{
			MethodName: "SayHello",
			Handler:    _DemomodSvc_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EchoStream",
			Handler:       _DemomodSvc_EchoStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "demomod.proto",
}