# Example
Middleware & bootstrapper designed for [gogf/gf](https://github.com/gogf/gf) web framework. 

## Documentation
- [Github](https://github.com/rookie-ninja/rk-gin)
- [Official Docs](https://docs.rkdev.info)

## Installation
- rk-boot: Bootstrapper base
- rk-gf: Bootstrapper for [gogf/gf](https://github.com/gogf/gf)

```shell
go get github.com/rookie-ninja/rk-boot/v2
go get github.com/rookie-ninja/rk-gf
```

## Quick start
### 1.Create boot.yaml
```yaml
---
# boot.yaml
gf:
  - name: rk-demo
    port: 8080
    enabled: true
    commonService:
      enabled: true
    sw:
      enabled: true
    docs:
      enabled: true
    prom:
      enabled: true
    middleware:
      logging:
        enabled: true
      prom:
        enabled: true
```

### 2.Create main.go
```go
// Copyright (c) 2021 rookie-ninja
//
// Use of this source code is governed by an Apache-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/rookie-ninja/rk-boot/v2"
	"github.com/rookie-ninja/rk-gf/boot"
	"net/http"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample rk-demo server.
// @termsOfService http://swagger.io/terms/

// @securityDefinitions.basic BasicAuth

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	// Create a new boot instance.
	boot := rkboot.NewBoot()

	// Register handler
	entry := rkgf.GetGfEntry("rk-demo")
	entry.Server.BindHandler("/v1/greeter", Greeter)

	// Bootstrap
	boot.Bootstrap(context.TODO())

	boot.WaitForShutdownSig(context.TODO())
}

// Greeter handler
// @Summary Greeter service
// @Id 1
// @version 1.0
// @produce application/json
// @Param name query string true "Input name"
// @Success 200 {object} GreeterResponse
// @Router /v1/greeter [get]
func Greeter(ctx *ghttp.Request) {
	ctx.Response.WriteHeader(http.StatusOK)
	ctx.Response.WriteJson(&GreeterResponse{
		Message: fmt.Sprintf("Hello %s!", ctx.GetQuery("name").String()),
	})
}

type GreeterResponse struct {
	Message string
}
```

### 3.Start server

```go
$ go run main.go
```

### 4.Validation
- Call API:

```shell script
$ curl -X GET localhost:8080/v1/greeter?name=rk-dev
{"Message":"Hello rk-dev!"}

$ curl -X GET localhost:8080/rk/v1/ready
{
  "ready": true
}

$ curl -X GET localhost:8080/rk/v1/alive
{
  "alive": true
}
```

- Swagger UI: [http://localhost:8080/sw](http://localhost:8080/sw)

- Docs UI via: [http://localhost:8080/docs](http://localhost:8080/docs)

- Prometheus client: [http://localhost:8080/metrics](http://localhost:8080/metrics)

