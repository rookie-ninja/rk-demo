# Simple GRPC server demo
This is the simplest grpc server demo with bellow functionality enabled.
- GRPC Server with server reflection
- With grpc-gateway

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [Quick start](#quick-start)
  - [Start server](#start-server)
  - [Send request](#send-request)
  - [Log output](#log-output)
  - [Directory layout](#directory-layout)
  - [boot.yaml](#bootyaml)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Quick start
### Start server
Run main.go in the terminal or run it from your IDE directly.

```go
go run main.go 
```

### Send request

```shell script
$ curl -X POST -d '{"name":"rk-dev"}' localhost:8080/api/v1/SayHello
{"message":"Hello rk-dev"}
```

### Log output
It will print logs something like bellow:
- The first line printed from default rkentry.ZapLoggerEntry which is commonly used logger format.
- The rest of logs are human readable format of rkentry.EventLoggerEntry which is used to log every event.

```text
2021-05-29T01:44:55.675+0800    INFO    boot/gw_entry.go:416    Bootstrapping GwEntry.  {"entryName": "greeter", "entryType": "GwEntry", "grpcPort": 1949, "httpPort": 8080, "swEnabled": false, "tvEnabled": false, "promEnabled": false, "commonServiceEnabled": false, "clientTlsEnabled": false, "serverTlsEnabled": false}
------------------------------------------------------------------------
endTime=2021-05-29T01:44:55.676077+08:00
startTime=2021-05-29T01:44:55.675824+08:00
elapsedNano=253341
hostname=lark.local
timing={}
counter={}
pair={}
error={}
field={"clientTlsEnabled":false,"commonServiceEnabled":false,"entryName":"greeter","entryType":"GwEntry","grpcPort":1949,"httpPort":8080,"promEnabled":false,"serverTlsEnabled":false,"swEnabled":false,"tvEnabled":false}
remoteAddr=lark.local
appName=rkApp
appVersion=unknown
entryName=greeter
entryType=GwEntry
locale=unknown
operation=bootstrap
eventStatus=Ended
timezone=CST
os=darwin
arch=amd64
EOE
2021-05-29T01:44:55.676+0800    INFO    boot/grpc_entry.go:610  Bootstrapping grpcEntry.        {"entryName": "greeter", "entryType": "GrpcEntry", "grpcPort": 1949, "commonServiceEnabled": false, "tlsEnabled": false, "gwEnabled": true, "reflectionEnabled": true, "swEnabled": false, "tvEnabled": false, "promEnabled": false, "gwClientTlsEnabled": false, "gwServerTlsEnabled": false}
------------------------------------------------------------------------
endTime=2021-05-29T01:44:55.676366+08:00
startTime=2021-05-29T01:44:55.675808+08:00
elapsedNano=557853
hostname=lark.local
timing={}
counter={}
pair={}
error={}
field={"commonServiceEnabled":false,"entryName":"greeter","entryType":"GrpcEntry","grpcPort":1949,"gwClientTlsEnabled":false,"gwEnabled":true,"gwServerTlsEnabled":false,"promEnabled":false,"reflectionEnabled":true,"swEnabled":false,"tlsEnabled":false,"tvEnabled":false}
remoteAddr=lark.local
appName=rkApp
appVersion=unknown
entryName=greeter
entryType=GrpcEntry
locale=unknown
operation=bootstrap
eventStatus=Ended
timezone=CST
os=darwin
arch=amd64
EOE
```

### Directory layout
simple-server demo contains 6 files in 2-server-with-gateway/ directory.

- boot.yaml
boot.yaml is the bootstrap config file for rk-boot, rk-boot will read this file to start GRPC server.
We locate boot.yaml file in the root working directory. As a result, we didn't specify file path of bootstrapper config file 
in the main.go function. Because rk-boot will looking for bootstrapper file in the root working directory named as boot.yaml
if not specified.

- main.go
Main entry for this program which includes implementation of grpc service described in proto file.

- api/
Contains proto files and gateway mapping config file.

- buf*.yaml
We use [buf](https://docs.buf.build/) to compile proto files. Please bellow command to recompile proto files which will whose output path 
is api/gen/ folder.

```shell script
make buf
```

```shell script
.
├── Makefile
├── README.md
├── api
│   ├── gen
│   │    └── v1
│   │        ├── hello.pb.go
│   │        ├── hello.pb.gw.go
│   │        ├── hello.swagger.json
│   │        └── hello_grpc.pb.go
│   └── v1
│        ├── gw_mapping.yaml
│        └── hello.proto
├── boot.yaml
├── buf.gen.yaml
├── buf.yaml
├── go.mod
├── go.sum
└── main.go
```

### main.go
```go
// Copyright (c) 2021 rookie-ninja
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
package main

import (
	"context"
	"github.com/rookie-ninja/rk-boot"
	"github.com/rookie-ninja/rk-demo/api/gen/v1"
	"google.golang.org/grpc"
)

func main() {
	// Create a new boot instance.
	boot := rkboot.NewBoot()

	// register grpc service
	boot.GetGrpcEntry("greeter").AddGrpcRegFuncs(registerGreeter)
    // register grpc-gateway handler
	boot.GetGrpcEntry("greeter").AddGwRegFuncs(greeter.RegisterGreeterHandlerFromEndpoint)

	// Bootstrap
	boot.Bootstrap(context.TODO())

	// Wait for shutdown signal
	boot.WaitForShutdownSig()

	// Interrupt entries
	boot.Interrupt(context.TODO())
}

func registerGreeter(server *grpc.Server) {
	greeter.RegisterGreeterServer(server, &GreeterServer{})
}

type GreeterServer struct{}

func (server *GreeterServer) SayHello(ctx context.Context, request *greeter.HelloRequest) (*greeter.HelloResponse, error) {
	return &greeter.HelloResponse{
		Message: "Hello " + request.Name,
	}, nil
}
```

### boot.yaml
We are using the simplest way of boot.yaml which contains only name and port which is required.

| name | description | type | default value |
| ------ | ------ | ------ | ------ |
| grpc.gw.enabled | Enable gateway service over gRpc server | boolean | false |
| grpc.gw.port | The port of gRpc gateway | integer | nil, server won't start |
| grpc.gw.gwMappingFilePaths | The grpc gateway mapping file path | string array | empty array |
| grpc.gw.cert.ref | Reference of cert entry declared in cert section | string | "" |

As a result, user will not obtain any of interceptors nor utility functions.

```yaml
---
---
grpc:
  - name: greeter             # Required
    port: 1949                # Required
    descirption: "demo grpc"  # Optional, default: ""
    reflection:               # Optional, default: false
      enabled: true
    gw:                       # Optional
      enabled: true           # Required
      port: 8080              # Required
```