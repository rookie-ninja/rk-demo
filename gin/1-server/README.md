# Simple Gin server demo
This is the simplest gin server demo with bellow functionality enabled.
- Gin Server

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
Send http request to localhost.

```go
$ curl -X POST -H "Content-Type: application/json" -d '{"name": "rk-dev"}' http://localhost:8080/v1/hello
{"response":"hello rk-dev"}
```

### Log output
It will print logs something like bellow:
- The first line printed from default rkentry.ZapLoggerEntry which is commonly used logger format.
- The rest of logs are human readable format of rkentry.EventLoggerEntry which is used to log every event.

```text
2021-05-26T18:17:39.519+0800    INFO    boot/gin_entry.go:553   Bootstrapping GinEntry. {"entryName": "greeter", "entryType": "GinEntry", "port": 8080, "interceptorsCount": 2, "swEnabled": false, "tlsEnabled": false, "commonServiceEnabled": false, "tvEnabled": false}
------------------------------------------------------------------------
endTime=2021-05-26T18:17:39.519929+08:00
startTime=2021-05-26T18:17:39.519812+08:00
elapsedNano=117779
hostname=lark.local
timing={}
counter={}
pair={}
error={}
field={"commonServiceEnabled":false,"entryName":"greeter","entryType":"GinEntry","interceptorsCount":2,"port":8080,"swEnabled":false,"tlsEnabled":false,"tvEnabled":false}
remoteAddr=lark.local
appName=rkApp
appVersion=unknown
entryName=greeter
entryType=GinEntry
locale=unknown
operation=bootstrap
eventStatus=Ended
timezone=CST
os=darwin
arch=amd64
EOE
```

### Directory layout
simple-server demo contains 6 files in simple-server/ directory.

boot.yaml is the bootstrap config file for rk-boot, rk-boot will read this file to start Gin server.
We locate boot.yaml file in the root working directory. As a result, we didn't specify file path of bootstrapper config file 
in the main.go function. Because rk-boot will looking for bootstrapper file in the root working directory named as boot.yaml
if not specified.

main.go is entry of main program which contains /v1/hello handler.

```shell script
.
├── Makefile
├── README.md
├── boot.yaml
├── go.mod
├── go.sum
└── main.go
```

### boot.yaml
We are using the simplest way of boot.yaml which contains only name and port which is required.

| name | description | type | default value |
| ------ | ------ | ------ | ------ |
| gin.name | Name of gin server entry | string | N/A |
| gin.port | Port of server | integer | nil, server won't start |
| gin.description | Description of server | string | "" |

As a result, user will not obtain any of interceptors nor utility functions.

```yaml
---
gin:
  - name: greeter                     # Required
    port: 8080                        # Required
```