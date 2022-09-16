package main

import (
	"context"
	_ "embed"
	"github.com/hibiken/asynq"
	"github.com/rookie-ninja/rk-boot/v2"
	"github.com/rookie-ninja/rk-demo/api/gen/v1"
	mytask "github.com/rookie-ninja/rk-demo/task"
	"github.com/rookie-ninja/rk-grpc/v2/boot"
	"github.com/rookie-ninja/rk-grpc/v2/middleware/context"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/grpc"
	"net/http"
)

//go:embed grpc-server.yaml
var bootYAML []byte

func main() {
	boot := rkboot.NewBoot(rkboot.WithBootConfigRaw(bootYAML))

	// register grpc
	grpcEntry := rkgrpc.GetGrpcEntry("greeter")
	grpcEntry.AddRegFuncGrpc(registerMyServer)
	grpcEntry.AddRegFuncGw(greeter.RegisterMyServerHandlerFromEndpoint)

	// Bootstrap
	boot.Bootstrap(context.TODO())

	// Wait for shutdown sig
	boot.WaitForShutdownSig(context.TODO())
}

func registerMyServer(server *grpc.Server) {
	greeter.RegisterMyServerServer(server, &MyServer{})
}

type MyServer struct{}

func (server *MyServer) Enqueue(ctx context.Context, req *greeter.TaskReq) (*greeter.TaskResp, error) {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "127.0.0.1:6379"})
	defer client.Close()

	// get trace metadata
	header := http.Header{}
	rkgrpcctx.GetTracerPropagator(ctx).Inject(ctx, propagation.HeaderCarrier(header))

	err := server.enqueueTask(client, header)
	err = server.enqueueTask(client, header)
	err = server.enqueueTask(client, header)

	return &greeter.TaskResp{}, err
}

func (server *MyServer) enqueueTask(client *asynq.Client, header http.Header) error {
	task, err := mytask.NewDemoTask(header)
	_, err = client.Enqueue(task)
	return err
}
