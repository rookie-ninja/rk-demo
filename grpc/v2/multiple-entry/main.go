// Copyright (c) 2021 rookie-ninja
//
// Use of this source code is governed by an Apache-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/rookie-ninja/rk-boot/v2"
	"github.com/rookie-ninja/rk-demo/api/gen/v1"
	"github.com/rookie-ninja/rk-gin/v2/boot"
	"github.com/rookie-ninja/rk-grpc/v2/boot"
	"google.golang.org/grpc"
	"net/http"
)

func main() {
	boot := rkboot.NewBoot()

	// register grpc
	grpcEntry := rkgrpc.GetGrpcEntry("grpc-api")
	grpcEntry.AddRegFuncGrpc(registerGreeter)
	grpcEntry.AddRegFuncGw(greeter.RegisterGreeterHandlerFromEndpoint)

	// register gin
	ginEntry := rkgin.GetGinEntry("rest-api")
	// we are using swagger config file located at api/gen/v1/greeter.swagger.json
	// please make sure path is defined as /v1/hello
	// this path should be as the same as path defined in api/v1/gw_mapping.yaml
	// since swagger UI will use this path to send request
	ginEntry.Router.Handle(http.MethodGet, "/v1/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Hello Gin user!")
	})

	// Bootstrap
	boot.Bootstrap(context.TODO())

	// Wait for shutdown sig
	boot.WaitForShutdownSig(context.TODO())
}

func registerGreeter(server *grpc.Server) {
	greeter.RegisterGreeterServer(server, &GreeterServer{})
}

//GreeterServer GreeterServer struct
type GreeterServer struct{}

// Hello response with hello message
func (server *GreeterServer) Hello(_ context.Context, _ *greeter.HelloRequest) (*greeter.HelloResponse, error) {
	return &greeter.HelloResponse{
		Message: "Hello gRPC user!",
	}, nil
}
