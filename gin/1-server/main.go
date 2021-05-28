// Copyright (c) 2021 rookie-ninja
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
package main

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"github.com/rookie-ninja/rk-boot"
	"net/http"
)

func main() {
	// Create a new boot instance.
	boot := rkboot.NewBoot()

	// Get gin entry with name specified in boot.yaml file.
	boot.GetGinEntry("greeter").Router.POST("/v1/hello", hello)

	// Bootstrap
	boot.Bootstrap(context.TODO())

	// Wait for shutdown signal
	boot.WaitForShutdownSig()

	// Interrupt entries
	boot.Interrupt(context.TODO())
}

// /v1/hello handler.
func hello(ctx *gin.Context) {
	ctx.Header("request-id", uuid.New().String())

	request := helloRequest{}

	if err := ctx.ShouldBindBodyWith(&request, binding.JSON); err == nil {
		ctx.JSON(http.StatusOK, &helloResponse{
			Response: "hello " + request.Name,
		})
	} else {
		NewError(ctx, http.StatusBadRequest, errors.New(err.Error()))
	}
}

type helloRequest struct {
	Name string `json:"name" yaml:"name" example:"rk-dev"`
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
