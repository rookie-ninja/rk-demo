package mid

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/rookie-ninja/rk-entry/v2/middleware/tracing"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"gopkg.in/yaml.v3"
	"net/http"
)

var (
	noopTracerProvider = trace.NewNoopTracerProvider()
)

const (
	spanKey    = "SpanKey"
	traceIdKey = "traceIdKey"
)

type BasePayload struct {
	Header http.Header `json:"Header"`
}

func NewTraceMiddleware(jaegerRaw []byte) (asynq.MiddlewareFunc, error) {
	conf := &rkmidtrace.BootConfig{}
	err := yaml.Unmarshal(jaegerRaw, conf)

	if err != nil {
		return nil, err
	}

	mid := &TraceMiddleware{
		set: rkmidtrace.NewOptionSet(rkmidtrace.ToOptions(conf, "worker", "ansynq")...),
	}

	return mid.Middleware, nil
}

type TraceMiddleware struct {
	set rkmidtrace.OptionSetInterface
}

func (m *TraceMiddleware) Middleware(h asynq.Handler) asynq.Handler {
	return asynq.HandlerFunc(func(ctx context.Context, t *asynq.Task) error {
		var p BasePayload
		if err := json.Unmarshal(t.Payload(), &p); err != nil {
			return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
		}

		ctx = m.set.GetPropagator().Extract(ctx, propagation.HeaderCarrier(p.Header))

		// create new span
		ctx, span := m.set.GetTracer().Start(ctx, t.Type())
		defer span.End()

		ctx = context.WithValue(ctx, spanKey, span)
		ctx = context.WithValue(ctx, traceIdKey, span.SpanContext().TraceID())

		err := h.ProcessTask(ctx, t)

		if err != nil {
			span.SetStatus(codes.Error, fmt.Sprintf("%v", err))
		} else {
			span.SetStatus(codes.Ok, "success")
		}

		return err
	})
}

func GetSpan(ctx context.Context) trace.Span {
	if v := ctx.Value(spanKey); v != nil {
		if res, ok := v.(trace.Span); ok {
			return res
		}
	}

	_, span := noopTracerProvider.Tracer("rk-trace-noop").Start(ctx, "noop-span")
	return span
}

func GetTraceId(ctx context.Context) string {
	return GetSpan(ctx).SpanContext().TraceID().String()
}
