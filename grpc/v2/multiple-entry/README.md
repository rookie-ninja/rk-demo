# Example
In this example, we will start Gin and gRPC in same process which shares same swagger config file.

```shell
go get github.com/rookie-ninja/rk-boot/v2
go get github.com/rookie-ninja/rk-grpc/v2
go get github.com/rookie-ninja/rk-gin/v2
```

## Quick start
### 0.Add dependency in third-party/ folder
Please add dependency files in third-party folder as example.

### 1.Prepare .proto files
- api/v1/greeter.proto

```protobuf
syntax = "proto3";

package api.v1;

option go_package = "api/v1/greeter";

service Greeter {
  rpc Hello (HelloRequest) returns (HelloResponse) {}
}

message HelloRequest {}

message HelloResponse {
  string message = 1;
}
```

- api/v1/gw_mapping.yaml

```yaml
type: google.api.Service
config_version: 3

# Please refer google.api.Http in third-party/googleapis/google/api/http.proto file for details.
http:
  rules:
    - selector: api.v1.Greeter.Hello
      get: /v1/hello
```

- buf.yaml

```yaml
version: v1beta1
name: github.com/rk-dev/rk-boot
build:
  roots:
    - api
    - third-party
```

- buf.gen.yaml

```yaml
version: v1beta1
plugins:
  - name: go
    out: api/gen
    opt:
     - paths=source_relative
  - name: go-grpc
    out: api/gen
    opt:
      - paths=source_relative
      - require_unimplemented_servers=false
  - name: grpc-gateway
    out: api/gen
    opt:
      - paths=source_relative
      - grpc_api_configuration=api/v1/gw_mapping.yaml
      - allow_repeated_fields_in_body=true
      - generate_unbound_methods=true
  - name: openapiv2
    out: api/gen
    opt:
      - grpc_api_configuration=api/v1/gw_mapping.yaml
      - allow_repeated_fields_in_body=true
```

### 2.Generate .pb.go files with [buf](https://docs.buf.build/introduction)
```
$ make buf
```

### 3.Create boot.yaml
Important note: rk-boot will bind grpc and grpc-gateway in the same port which we think is a convenient way.

As a result, grpc-gateway will automatically be started.

```yaml
---
grpc:
  - name: grpc-api
    enabled: true
    port: 8080
    enableReflection: true
    sw:
      enabled: true
      # define swagger config file path if we hope to use different one
      # by default, rk-boot will looking for swagger config file from api/gen/v1 and docs/ folder
      # jsonPath: ""
gin:
  - name: rest-api
    port: 8081
    enabled: true
    sw:
      enabled: true
      # define swagger config file path if we hope to use different one
      # by default, rk-boot will looking for swagger config file from api/gen/v1 and docs/ folder
      # jsonPath: ""
```

### 4.Create main.go
```go
// Copyright (c) 2021 rookie-ninja
//
// Use of this source code is governed by an Apache-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/rookie-ninja/rk-boot/v2"
	"github.com/rookie-ninja/rk-demo/api/gen/v1"
	"github.com/rookie-ninja/rk-gin/v2/boot"
	"github.com/rookie-ninja/rk-grpc/v2/boot"
	"google.golang.org/grpc"
	"net/http"
)

func main() {
	boot := rkboot.NewBoot()

	// register grpc
	grpcEntry := rkgrpc.GetGrpcEntry("grpc-api")
	grpcEntry.AddRegFuncGrpc(registerGreeter)
	grpcEntry.AddRegFuncGw(greeter.RegisterGreeterHandlerFromEndpoint)

	// register gin
	ginEntry := rkgin.GetGinEntry("rest-api")
	// we are using swagger config file located at api/gen/v1/greeter.swagger.json
	// please make sure path is defined as /v1/hello
	// this path should be as the same as path defined in api/v1/gw_mapping.yaml
	// since swagger UI will use this path to send request
	ginEntry.Router.Handle(http.MethodGet, "/v1/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Hello Gin user!")
	})

	// Bootstrap
	boot.Bootstrap(context.TODO())

	// Wait for shutdown sig
	boot.WaitForShutdownSig(context.TODO())
}

func registerGreeter(server *grpc.Server) {
	greeter.RegisterGreeterServer(server, &GreeterServer{})
}

//GreeterServer GreeterServer struct
type GreeterServer struct{}

// Hello response with hello message
func (server *GreeterServer) Hello(_ context.Context, _ *greeter.HelloRequest) (*greeter.HelloResponse, error) {
	return &greeter.HelloResponse{
		Message: "Hello gRPC user!",
	}, nil
}
```

### 5.Start server

```go
$ go run main.go
```

### 4.Validation
#### 4.1 Validate Gin server
- Swagger UI: [http://localhost:8081/sw](http://localhost:8081/sw)

```shell
$ curl localhost:8081/v1/hello
"Hello Gin user!"
```

#### 4.2 Validate gRPC server
- Swagger UI: [http://localhost:8080/sw](http://localhost:8080/sw)

Restful API with grpc-gateway which automatically started!

```shell
$ curl localhost:8080/v1/hello
{"message":"Hello gRPC user!"}
```

gRPC protocol with grpcurl

```shell
$ grpcurl -plaintext localhost:8080 api.v1.Greeter.Hello 
{
  "message": "Hello gRPC user!"
}
```