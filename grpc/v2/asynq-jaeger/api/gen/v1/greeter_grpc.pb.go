// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package greeter

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

// MyServerClient is the client API for MyServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MyServerClient interface {
	Enqueue(ctx context.Context, in *TaskReq, opts ...grpc.CallOption) (*TaskResp, error)
}

type myServerClient struct {
	cc grpc.ClientConnInterface
}

func NewMyServerClient(cc grpc.ClientConnInterface) MyServerClient {
	return &myServerClient{cc}
}

func (c *myServerClient) Enqueue(ctx context.Context, in *TaskReq, opts ...grpc.CallOption) (*TaskResp, error) {
	out := new(TaskResp)
	err := c.cc.Invoke(ctx, "/api.v1.MyServer/Enqueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MyServerServer is the server API for MyServer service.
// All implementations should embed UnimplementedMyServerServer
// for forward compatibility
type MyServerServer interface {
	Enqueue(context.Context, *TaskReq) (*TaskResp, error)
}

// UnimplementedMyServerServer should be embedded to have forward compatible implementations.
type UnimplementedMyServerServer struct {
}

func (UnimplementedMyServerServer) Enqueue(context.Context, *TaskReq) (*TaskResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enqueue not implemented")
}

// UnsafeMyServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MyServerServer will
// result in compilation errors.
type UnsafeMyServerServer interface {
	mustEmbedUnimplementedMyServerServer()
}

func RegisterMyServerServer(s grpc.ServiceRegistrar, srv MyServerServer) {
	s.RegisterService(&MyServer_ServiceDesc, srv)
}

func _MyServer_Enqueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyServerServer).Enqueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.MyServer/Enqueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyServerServer).Enqueue(ctx, req.(*TaskReq))
	}
	return interceptor(ctx, in, info, handler)
}

// MyServer_ServiceDesc is the grpc.ServiceDesc for MyServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MyServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.MyServer",
	HandlerType: (*MyServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Enqueue",
			Handler:    _MyServer_Enqueue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/greeter.proto",
}
