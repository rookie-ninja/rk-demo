// Copyright (c) 2021 rookie-ninja
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
package main

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rookie-ninja/rk-boot"
	rkprom "github.com/rookie-ninja/rk-prom"
	"net/http"
)

// @title Swagger Example API
// @version 1.0
// @description This is gin server with rk-boot.
// @termsOfService http://swagger.io/terms/
// @securityDefinitions.basic BasicAuth
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @schemes http https
func main() {
	// Create a new boot instance.
	boot := rkboot.NewBoot()

	// Register handler
	boot.GetGinEntry("greeter").Router.GET("/v1/hello", hello)

	// Bootstrap
	boot.Bootstrap(context.TODO())

	// ***************** Create custom metrics *****************
	// Since we enabled prometheus client, we could get PromEntry
	promEntry := boot.GetGinEntry("greeter").PromEntry

	// 1: Create a metrics set
	// MetricsSet is a collection of counter, gauge, summary, histogram metrics.
	// Each MetricsSet could be distinguished by <namespace> and <subsystem>.
	// <namespace> and <subsystem> should follow bellow regex which is required by prometheus:
	// ^[a-zA-Z_:][a-zA-Z0-9_:]*$
	metricsSet := rkprom.NewMetricsSet("my_ns", "my_sub", promEntry.Registerer)

	// 2: Register a metrics(counter, gauge, summary, histogram)
	// Register a counter with name and label.
	// Label should also follow regex of ^[a-zA-Z_:][a-zA-Z0-9_:]*$
	if err := metricsSet.RegisterCounter("my_counter", "label_1", "label_2"); err != nil {
		panic(err)
	}

	// 3: Record data into metrics
	// Get Counter with values and increase it.
	// Number of values and types should match labels we set already.
	metricsSet.GetCounterWithValues("my_counter", "value_1", "value_2").Inc()

	// Wait for shutdown signal
	boot.WaitForShutdownSig()

	// Interrupt entries
	boot.Interrupt(context.TODO())
}

// @Summary Hello
// @Description Say hello to incoming name.
// @Id v1.api.hello
// @Accept  application/json
// @Tags Hello
// @version 1.0
// @Param name query string true "Your name"
// @Produce application/json
// @Success 200 {object} helloResponse
// @Failure 400 {object} httpError
// @Router /v1/hello [get]
// @Header all {string} request-id "Request id for with uuid generator."
func hello(ctx *gin.Context) {
	ctx.Header("request-id", uuid.New().String())

	if name := ctx.Query("name"); len(name) < 1 {
		NewError(ctx, http.StatusBadRequest, errors.New("name should not be nil"))
		return
	}

	ctx.JSON(http.StatusOK, &helloResponse{
		Response: "hello " + ctx.Query("name"),
	})
}

type helloResponse struct {
	Response string `json:"response" yaml:"response" example:"hello user"`
}

func NewError(ctx *gin.Context, status int, err error) {
	ctx.JSON(status, httpError{
		Code:    status,
		Message: err.Error(),
	})
}

type httpError struct {
	Code    int    `json:"code" yaml:"code" example:"400"`
	Message string `json:"message" yaml:"message" example:"status bad request"`
}
