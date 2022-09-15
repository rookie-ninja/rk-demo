package main

import (
	"context"
	_ "embed"
	"github.com/hibiken/asynq"
	"github.com/rookie-ninja/rk-boot/v2"
	"github.com/rookie-ninja/rk-demo/api/gen/v1"
	mytask "github.com/rookie-ninja/rk-demo/task"
	"github.com/rookie-ninja/rk-grpc/v2/boot"
	rkgrpcctx "github.com/rookie-ninja/rk-grpc/v2/middleware/context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
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

	task, err := mytask.NewDemoTask(rkgrpcctx.GetTraceId(ctx))
	if err != nil {
		return nil, err
	}
	info, err := client.Enqueue(task)
	if err != nil {
		return nil, err
	}

	logger := rkgrpcctx.GetLogger(ctx)
	logger.Info("enqueued task", zap.String("id", info.ID), zap.String("queue", info.Queue))

	return &greeter.TaskResp{}, nil
}
