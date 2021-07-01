// Copyright (c) 2021 rookie-ninja
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
package main

import (
	"context"
	"github.com/rookie-ninja/rk-boot"
	"github.com/rookie-ninja/rk-demo/api/gen/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	// Create a new boot instance.
	boot := rkboot.NewBoot()

	// register grpc service
	boot.GetGrpcEntry("greeter").AddGrpcRegFuncs(registerGreeter)
	// register grpc-gateway handler
	boot.GetGrpcEntry("greeter").AddGwRegFuncs(greeter.RegisterGreeterHandlerFromEndpoint)

	// Bootstrap
	boot.Bootstrap(context.TODO())

	// Wait for shutdown signal
	boot.WaitForShutdownSig()

	// Interrupt entries
	boot.Interrupt(context.TODO())
}

func registerGreeter(server *grpc.Server) {
	greeter.RegisterGreeterServer(server, &GreeterServer{})
}

type GreeterServer struct{}

func (server *GreeterServer) SayHello(ctx context.Context, request *greeter.HelloRequest) (*greeter.HelloResponse, error) {
	if len(request.Name) < 1 {
		std := status.New(codes.InvalidArgument		, "name should not be empty!")
		std, _ = std.WithDetails(&greeter.HelloResponse{
			Message: "test",
		})

		return nil, std.Err()
	}

	return &greeter.HelloResponse{
		Message: "Hello " + request.Name,
	}, nil
}
