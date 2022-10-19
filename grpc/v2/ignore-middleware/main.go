// Copyright (c) 2021 rookie-ninja
//
// Use of this source code is governed by an Apache-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"github.com/rookie-ninja/rk-boot/v2"
	"github.com/rookie-ninja/rk-demo/api/gen/v1"
	"github.com/rookie-ninja/rk-grpc/v2/boot"
	"google.golang.org/grpc"
)

func main() {
	boot := rkboot.NewBoot()

	// register grpc
	grpcEntry := rkgrpc.GetGrpcEntry("greeter")
	grpcEntry.AddRegFuncGrpc(registerGreeter)
	grpcEntry.AddRegFuncGw(greeter.RegisterGreeterHandlerFromEndpoint)

	// Bootstrap
	boot.Bootstrap(context.TODO())

	// Wait for shutdown sig
	boot.WaitForShutdownSig(context.TODO())
}

func registerGreeter(server *grpc.Server) {
	greeter.RegisterGreeterServer(server, &GreeterServer{})
}

// GreeterServer GreeterServer struct
type GreeterServer struct{}

func (server *GreeterServer) Ignore(_ context.Context, _ *greeter.IgnoreRequest) (*greeter.IgnoreResponse, error) {
	return &greeter.IgnoreResponse{}, nil
}

func (server *GreeterServer) Show(_ context.Context, _ *greeter.ShowRequest) (*greeter.ShowResponse, error) {
	return &greeter.ShowResponse{}, nil
}
