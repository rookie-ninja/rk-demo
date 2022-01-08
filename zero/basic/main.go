// Copyright (c) 2021 rookie-ninja
//
// Use of this source code is governed by an Apache-style
// license that can be found in the LICENSE file.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/rookie-ninja/rk-boot"
	"github.com/rookie-ninja/rk-boot/zero"
	"github.com/tal-tech/go-zero/rest"
	"net/http"
)

// @title RK Swagger for Zero
// @version 1.0
// @description This is a greeter service with rk-boot.

// Application entrance.
func main() {
	// Create a new boot instance.
	boot := rkboot.NewBoot()

	// Register handler
	zeroEntry := rkbootzero.GetZeroEntry("greeter")
	zeroEntry.Server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/v1/greeter",
		Handler: Greeter,
	})

	// Bootstrap
	boot.Bootstrap(context.Background())

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
func Greeter(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	bytes, _ := json.Marshal(&GreeterResponse{
		Message: fmt.Sprintf("Hello %s!", request.URL.Query().Get("name")),
	})

	writer.Write(bytes)
}

// Response.
type GreeterResponse struct {
	Message string
}
