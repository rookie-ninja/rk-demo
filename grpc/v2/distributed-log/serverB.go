package main

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/rookie-ninja/rk-boot/v2"
	"github.com/rookie-ninja/rk-demo/api/gen/v1"
	"github.com/rookie-ninja/rk-grpc/v2/boot"
	"google.golang.org/grpc"
)

//go:embed bootB.yaml
var bootB []byte

func main() {
	boot := rkboot.NewBoot(rkboot.WithBootConfigRaw(bootB))

	// register grpc
	grpcEntry := rkgrpc.GetGrpcEntry("serverB")
	grpcEntry.AddRegFuncGrpc(registerGreeterB)
	grpcEntry.AddRegFuncGw(greeter.RegisterGreeterHandlerFromEndpoint)

	// Bootstrap
	boot.Bootstrap(context.TODO())

	// Wait for shutdown sig
	boot.WaitForShutdownSig(context.TODO())
}

func registerGreeterB(server *grpc.Server) {
	greeter.RegisterGreeterServer(server, &GreeterServerB{})
}

type GreeterServerB struct{}

// Hello response with hello message
func (server *GreeterServerB) Hello(ctx context.Context, req *greeter.HelloRequest) (*greeter.HelloResponse, error) {
	return &greeter.HelloResponse{
		Message: fmt.Sprintf("Hello %s!", req.Name),
	}, nil
}
