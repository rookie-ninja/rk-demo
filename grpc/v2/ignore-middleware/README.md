# Example
In this example, we will configure boot.yaml to ignore middleware.

```shell
go get github.com/rookie-ninja/rk-boot/v2
go get github.com/rookie-ninja/rk-grpc/v2
```

## Quick start
### 1.Create boot.yaml
Important note: rk-boot will bind grpc and grpc-gateway in the same port which we think is a convenient way.

As a result, grpc-gateway will automatically be started.

```yaml
---
grpc:
  - name: greeter
    enabled: true
    port: 8080
    enableReflection: true
    enableRkGwOption: true
    middleware:
      trace:
        enabled: true
        ignore: ["/api.v1.Greeter/Ignore"]
        exporter:
          file:
            enabled: true
            outputPath: "stdout"
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
	"github.com/rookie-ninja/rk-boot/v2"
	"github.com/rookie-ninja/rk-demo/api/gen/v1"
	"github.com/rookie-ninja/rk-grpc/v2/boot"
	"google.golang.org/grpc"
)

func main() {
	boot := rkboot.NewBoot()

	// register grpc
	grpcEntry := rkgrpc.GetGrpcEntry("greeter")
	grpcEntry.AddRegFuncGrpc(registerGreeter)
	grpcEntry.AddRegFuncGw(greeter.RegisterGreeterHandlerFromEndpoint)

	// Bootstrap
	boot.Bootstrap(context.TODO())

	// Wait for shutdown sig
	boot.WaitForShutdownSig(context.TODO())
}

func registerGreeter(server *grpc.Server) {
	greeter.RegisterGreeterServer(server, &GreeterServer{})
}

// GreeterServer GreeterServer struct
type GreeterServer struct{}

func (server *GreeterServer) Ignore(_ context.Context, _ *greeter.IgnoreRequest) (*greeter.IgnoreResponse, error) {
	return &greeter.IgnoreResponse{}, nil
}

func (server *GreeterServer) Show(_ context.Context, _ *greeter.ShowRequest) (*greeter.ShowResponse, error) {
	return &greeter.ShowResponse{}, nil
}

```

### 3.Start server

```go
$ go run main.go
```

### 4.Validation
#### 4.1 Ignore trace
No trace logs expected.

```shell
$ curl localhost:8080/v1/ignore
```

#### 4.1 Show trace

```shell
$ curl localhost:8080/v1/show

{
        "Name": "/api.v1.Greeter/Show",
        "SpanContext": {
                "TraceID": "e45c21793bb5a8359b844e322ecb8270",
                "SpanID": "f659f335f4b70065",
                "TraceFlags": "01",
                "TraceState": "",
                "Remote": false
        },
        "Parent": {
                "TraceID": "00000000000000000000000000000000",
                "SpanID": "0000000000000000",
                "TraceFlags": "00",
                "TraceState": "",
                "Remote": true
        },
        "SpanKind": 2,
...
```