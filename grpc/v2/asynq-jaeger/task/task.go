package task

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/rookie-ninja/rk-entry/v2/entry"
	"github.com/rookie-ninja/rk-repo/asynq"
	"go.uber.org/zap"
	"net/http"
	"time"
)

// ************ Task ************

const (
	TypeDemo = "demo-task"
)

type DemoPayload struct {
	TraceHeader http.Header `json:"traceHeader"`
}

func NewDemoTask(header http.Header) (*asynq.Task, error) {
	payload, err := json.Marshal(DemoPayload{
		TraceHeader: header,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeDemo, payload), nil
}

func HandleDemoTask(ctx context.Context, t *asynq.Task) error {
	// sleep a while for testing
	time.Sleep(50 * time.Millisecond)

	var p DemoPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	rkentry.GlobalAppCtx.GetLoggerEntryDefault().Info("handle demo task", zap.String("traceId", rkasynq.GetTraceId(ctx)))

	CallFuncA(ctx)
	CallFuncB(ctx)

	return nil
}

func CallFuncA(ctx context.Context) {
	newCtx, span := rkasynq.NewSpan(ctx, "funcA")
	defer rkasynq.EndSpan(span, true)

	time.Sleep(10 * time.Millisecond)

	CallFuncAA(newCtx)
}

func CallFuncAA(ctx context.Context) {
	_, span := rkasynq.NewSpan(ctx, "funcAA")
	defer rkasynq.EndSpan(span, true)

	time.Sleep(10 * time.Millisecond)
}

func CallFuncB(ctx context.Context) {
	_, span := rkasynq.NewSpan(ctx, "funcB")
	defer rkasynq.EndSpan(span, true)

	time.Sleep(10 * time.Millisecond)
}