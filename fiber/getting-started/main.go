// Copyright (c) 2021 rookie-ninja
//
// Use of this source code is governed by an Apache-style
// license that can be found in the LICENSE file.
package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rookie-ninja/rk-boot"
	"github.com/rookie-ninja/rk-boot/fiber"
	"net/http"
)

// @title RK Swagger for Fiber
// @version 1.0
// @description This is a greeter service with rk-boot.

// Application entrance.
func main() {
	// Create a new boot instance.
	boot := rkboot.NewBoot()

	// Bootstrap
	boot.Bootstrap(context.Background())

	// Register handler
	// Because middleware needs to be added into *fiber.App before any router,
	// We need to init router after Bootstrap() function.
	//
	// Otherwise, middleware won't work!
	fiberEntry := rkbootfiber.GetFiberEntry("greeter")
	fiberEntry.App.Get("/v1/greeter", Greeter)

	// Wait for shutdown sig
	boot.WaitForShutdownSig(context.Background())
}

// @Summary Greeter service
// @Id 1
// @version 1.0
// @produce application/json
// @Param name query string true "Input name"
// @Success 200 {object} GreeterResponse
// @Router /v1/greeter [get]
func Greeter(ctx *fiber.Ctx) error {
	ctx.Response().SetStatusCode(http.StatusOK)
	return ctx.JSON(&GreeterResponse{
		Message: fmt.Sprintf("Hello %s!", ctx.Query("name")),
	})
}

// Response.
type GreeterResponse struct {
	Message string
}
