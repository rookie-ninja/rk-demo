package task

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/rookie-ninja/rk-demo/mid"
	rkentry "github.com/rookie-ninja/rk-entry/v2/entry"
	"go.uber.org/zap"
	"net/http"
)

// ************ Task ************

const (
	TypeDemo = "demo:test"
)

type DemoPayload struct {
	Header http.Header `json:"Header"`
}

func NewDemoTask(header http.Header) (*asynq.Task, error) {
	payload, err := json.Marshal(DemoPayload{
		Header: header,
	})
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

	rkentry.GlobalAppCtx.GetLoggerEntryDefault().Info("handle demo task", zap.String("traceId", mid.GetTraceId(ctx)))

	return nil
}
