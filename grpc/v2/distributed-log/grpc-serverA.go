package main

import (
	"context"
	_ "embed"
	"github.com/rookie-ninja/rk-boot/v2"
	greeter "github.com/rookie-ninja/rk-demo/api/gen/v1"
	"github.com/rookie-ninja/rk-grpc/v2/boot"
	"google.golang.org/grpc"
)

//go:embed grpc-serverA.yaml
var grpcBootA []byte

func main() {
	boot := rkboot.NewBoot(rkboot.WithBootConfigRaw(grpcBootA))

	// register grpc
	grpcEntry := rkgrpc.GetGrpcEntry("grpcServerA")
	grpcEntry.AddRegFuncGrpc(registerServerA)

	// Bootstrap
	boot.Bootstrap(context.TODO())

	// Wait for shutdown sig
	boot.WaitForShutdownSig(context.TODO())
}

func registerServerA(server *grpc.Server) {
	greeter.RegisterServerAServer(server, &ServerA{})
}

type ServerA struct{}

func (server *ServerA) Login(ctx context.Context, req *greeter.LoginReq) (*greeter.LoginResp, error) {
	return &greeter.LoginResp{}, nil
}

func (server *ServerA) CallA(ctx context.Context, req *greeter.CallAReq) (*greeter.CallAResp, error) {
	return &greeter.CallAResp{}, nil
}
