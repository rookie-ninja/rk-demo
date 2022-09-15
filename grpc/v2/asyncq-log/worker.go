package main

import (
	_ "embed"
	"github.com/hibiken/asynq"
	"github.com/rookie-ninja/rk-demo/mid"
	"github.com/rookie-ninja/rk-demo/task"
	"github.com/rookie-ninja/rk-entry/v2/entry"
)

var logger = rkentry.GlobalAppCtx.GetLoggerEntryDefault().Logger

//go:embed jaeger.yaml
var jaegerYaml []byte

func main() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: "127.0.0.1:6379"},
		asynq.Config{
			Logger: logger.Sugar(),
		},
	)

	// mux maps a type to a handler
	mux := asynq.NewServeMux()
	mux.HandleFunc(task.TypeDemo, task.HandleDemoTask)

	traceMid, err := mid.NewTraceMiddleware(jaegerYaml)
	if err != nil {
		rkentry.ShutdownWithError(err)
	}
	mux.Use(traceMid)

	if err := srv.Run(mux); err != nil {
		rkentry.ShutdownWithError(err)
	}
}
