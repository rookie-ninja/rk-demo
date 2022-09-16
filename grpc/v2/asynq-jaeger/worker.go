package main

import (
	_ "embed"
	"github.com/hibiken/asynq"
	"github.com/rookie-ninja/rk-demo/task"
	"github.com/rookie-ninja/rk-entry/v2/entry"
	"github.com/rookie-ninja/rk-repo/asynq"
)

var logger = rkentry.GlobalAppCtx.GetLoggerEntryDefault().Logger

//go:embed trace.yaml
var traceConfRaw []byte

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

	// add jaeger middleware
	jaegerMid, err := rkasynq.NewJaegerMid(traceConfRaw)
	if err != nil {
		rkentry.ShutdownWithError(err)
	}
	mux.Use(jaegerMid)

	if err := srv.Run(mux); err != nil {
		rkentry.ShutdownWithError(err)
	}
}
