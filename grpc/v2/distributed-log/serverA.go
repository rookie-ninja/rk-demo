package main

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/rookie-ninja/rk-boot/v2"
	"github.com/rookie-ninja/rk-demo/api/gen/v1"
	"github.com/rookie-ninja/rk-grpc/v2/boot"
	"github.com/rookie-ninja/rk-grpc/v2/middleware/context"
	"google.golang.org/grpc"
)

//go:embed bootA.yaml
var bootA []byte

func main() {
	boot := rkboot.NewBoot(rkboot.WithBootConfigRaw(bootA))

	// register grpc
	grpcEntry := rkgrpc.GetGrpcEntry("serverA")
	grpcEntry.AddRegFuncGrpc(registerGreeterA)
	grpcEntry.AddRegFuncGw(greeter.RegisterGreeterHandlerFromEndpoint)

	// Bootstrap
	boot.Bootstrap(context.TODO())

	// Wait for shutdown sig
	boot.WaitForShutdownSig(context.TODO())
}

func registerGreeterA(server *grpc.Server) {
	greeter.RegisterGreeterServer(server, &GreeterServerA{})
}

type GreeterServerA struct{}

// Hello response with hello message
func (server *GreeterServerA) Hello(ctx context.Context, req *greeter.HelloRequest) (*greeter.HelloResponse, error) {
	// Call serverB at 2008 with grpc client
	opts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithInsecure(),
	}
	conn, _ := grpc.Dial("localhost:2008", opts...)
	defer conn.Close()
	client := greeter.NewGreeterClient(conn)

	// Inject current trace information into context
	newCtx := rkgrpcctx.InjectSpanToNewContext(ctx)
	client.Hello(newCtx, &greeter.HelloRequest{Name: "A"})

	return &greeter.HelloResponse{
		Message: fmt.Sprintf("Hello %s!", req.Name),
	}, nil
}
