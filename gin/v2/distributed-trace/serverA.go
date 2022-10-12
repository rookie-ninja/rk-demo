// Copyright (c) 2021 rookie-ninja
//
// Use of this source code is governed by an Apache-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rookie-ninja/rk-boot/v2"
	"github.com/rookie-ninja/rk-gin/v2/boot"
	"github.com/rookie-ninja/rk-gin/v2/middleware/context"
	"github.com/valyala/fasthttp"
	"go.opentelemetry.io/otel/trace"
	"net/http"
)

//go:embed serverA.yaml
var rawBootA []byte

func main() {
	// Create a new boot instance.
	boot := rkboot.NewBoot(rkboot.WithBootConfigRaw(rawBootA))

	// Register handler
	entry := rkgin.GetGinEntry("serverA")
	entry.Router.GET("/v1/first", First)

	// Bootstrap
	boot.Bootstrap(context.TODO())

	boot.WaitForShutdownSig(context.TODO())
}

func First(ctx *gin.Context) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI("http://localhost:2008/v1/second")
	req.Header.SetMethod(http.MethodGet)

	InjectSpanToRequest(ctx, req)

	resp := fasthttp.AcquireResponse()

	fasthttp.Do(req, resp)

	bodyBytes := resp.Body()
	println(string(bodyBytes))

	ctx.JSON(http.StatusOK, &FirstResponse{
		Message: "First response",
	})
}

type FirstResponse struct {
	Message string
}

// *************** For Trace of fasthttp ***************

// InjectSpanToRequest inject trace span into request
func InjectSpanToRequest(ctx *gin.Context, req *fasthttp.Request) {
	newCtx := trace.ContextWithRemoteSpanContext(context.Background(), rkginctx.GetTraceSpan(ctx).SpanContext())
	if propagator := rkginctx.GetTracerPropagator(ctx); propagator != nil {
		propagator.Inject(newCtx, &FastHttpHeaderCarrier{Req: req})
	}
}

// FastHttpHeaderCarrier a new carrier needs to be implemented
type FastHttpHeaderCarrier struct {
	Req *fasthttp.Request
}

// Get returns the value associated with the passed key.
func (hc *FastHttpHeaderCarrier) Get(key string) string {
	return string(hc.Req.Header.Peek(key))
}

// Set stores the key-value pair.
func (hc *FastHttpHeaderCarrier) Set(key string, value string) {
	hc.Req.Header.Set(key, value)
}

// Keys lists the keys stored in this carrier.
func (hc *FastHttpHeaderCarrier) Keys() []string {
	res := make([]string, 0)
	hc.Req.Header.VisitAll(func(key, value []byte) {
		res = append(res, fmt.Sprintf("%v", key))
	})
	return res
}
