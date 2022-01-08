// Copyright (c) 2021 rookie-ninja
//
// Use of this source code is governed by an Apache-style
// license that can be found in the LICENSE file.
package main

import (
	"context"
	"github.com/rookie-ninja/rk-boot"
	"github.com/rookie-ninja/rk-boot/fiber"
	"github.com/rookie-ninja/rk-demo/internal/api/v1"
)

// @title RK Swagger for Fiber
// @version 1.0
// @description This is a greeter service with rk-boot.
// @termsOfService http://swagger.io/terms/

// @securityDefinitions.basic BasicAuth

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @name Authorization

// @schemes http https

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
	fiberEntry.App.Get("/v1/greeter", api.Greeter)

	// Wait for shutdown sig
	boot.WaitForShutdownSig(context.Background())
}
