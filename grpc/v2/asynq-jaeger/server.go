package main

import (
	"context"
	_ "embed"
	"github.com/hibiken/asynq"
	"github.com/rookie-ninja/rk-boot/v2"
	"github.com/rookie-ninja/rk-demo/api/gen/v1"
	"github.com/rookie-ninja/rk-demo/task"
	"github.com/rookie-ninja/rk-entry/v2/entry"
	"github.com/rookie-ninja/rk-grpc/v2/boot"
	"github.com/rookie-ninja/rk-grpc/v2/middleware/context"
	"github.com/rookie-ninja/rk-repo/asynq"
	"go.opentelemetry.io/otel/propagation"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net/http"
)

//go:embed boot.yaml
var bootYAML []byte

func main() {
	boot := rkboot.NewBoot(rkboot.WithBootConfigRaw(bootYAML))

	// register grpc
	grpcEntry := rkgrpc.GetGrpcEntry("greeter")
	grpcEntry.AddRegFuncGrpc(registerMyServer)
	grpcEntry.AddRegFuncGw(greeter.RegisterMyServerHandlerFromEndpoint)

	// Bootstrap
	boot.Bootstrap(context.TODO())

	asynqServer := startAsynqWorker(grpcEntry.LoggerEntry.Logger)
	boot.AddShutdownHookFunc("asynq worker", asynqServer.Stop)

	// Wait for shutdown sig
	boot.WaitForShutdownSig(context.TODO())
}

func startAsynqWorker(logger *zap.Logger) *asynq.Server {
	// start asynq service
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: "127.0.0.1:6379"},
		asynq.Config{
			Logger: logger.Sugar(),
		},
	)

	// mux maps a type to a handler
	mux := asynq.NewServeMux()
	mux.HandleFunc(task.TypeDemo, task.HandleDemoTask)

	// add jaeger middleware
	jaegerMid, err := rkasynq.NewJaegerMid(bootYAML)
	if err != nil {
		rkentry.ShutdownWithError(err)
	}
	mux.Use(jaegerMid)

	if err := srv.Start(mux); err != nil {
		rkentry.ShutdownWithError(err)
	}

	return srv
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

	return &greeter.TaskResp{}, err
}

func (server *MyServer) enqueueTask(client *asynq.Client, header http.Header) error {
	task, err := task.NewDemoTask(header)
	_, err = client.Enqueue(task)
	return err
}
