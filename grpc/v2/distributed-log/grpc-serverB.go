package main

import (
	"context"
	_ "embed"
	"github.com/rookie-ninja/rk-boot/v2"
	greeter "github.com/rookie-ninja/rk-demo/api/gen/v1"
	"github.com/rookie-ninja/rk-grpc/v2/boot"
	"google.golang.org/grpc"
)

//go:embed grpc-serverB.yaml
var grpcBootB []byte

func main() {
	boot := rkboot.NewBoot(rkboot.WithBootConfigRaw(grpcBootB))

	// register grpc
	grpcEntry := rkgrpc.GetGrpcEntry("grpcServerB")
	grpcEntry.AddRegFuncGrpc(registerServerB)

	// Bootstrap
	boot.Bootstrap(context.TODO())

	// Wait for shutdown sig
	boot.WaitForShutdownSig(context.TODO())
}

func registerServerB(server *grpc.Server) {
	greeter.RegisterServerBServer(server, &ServerB{})
}

type ServerB struct{}

func (s ServerB) CallB(ctx context.Context, req *greeter.CallBReq) (*greeter.CallBResp, error) {
	return &greeter.CallBResp{}, nil
}
