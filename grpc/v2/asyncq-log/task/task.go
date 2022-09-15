package task

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/rookie-ninja/rk-entry/v2/entry"
	"go.uber.org/zap"
)

// ************ Task ************

const (
	TypeDemo = "demo:test"
)

type DemoPayload struct {
	TraceId string
}

func NewDemoTask(traceId string) (*asynq.Task, error) {
	payload, err := json.Marshal(DemoPayload{TraceId: traceId})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeDemo, payload), nil
}

func HandleDemoTask(ctx context.Context, t *asynq.Task) error {
	var p DemoPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	rkentry.GlobalAppCtx.GetLoggerEntryDefault().Info("handle demo task", zap.String("traceId", p.TraceId))

	return nil
}
