# Simple Gin server demo
This is the simplest gin server demo with bellow functionality enabled.
- Gin Server
- Swagger UI
- RK common service (A list of commonly used APIs)
- Prometheus client
- Logging interceptor
- Metrics interceptor
- BasicAuth interceptor
- RK TV
- Application Info
- Zap Logger

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [Quick start](#quick-start)
  - [Start server](#start-server)
  - [Check logs/ directory](#check-logs-directory)
  - [Directory layout](#directory-layout)
  - [boot.yaml](#bootyaml)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Quick start
### Start server
Run main.go in the terminal or run it from your IDE directly.

```go
go run main.go 
```

### Check logs/ directory
New greeter.log file would be created automatically.

```text
2021-05-27T13:16:47.269+0800	INFO	Bootstrapping SwEntry.	{"domain": "rk-demo", "entryName": "greeter-sw", "entryType": "GinSwEntry", "jsonPath": "docs", "path": "/sw/", "port": 8080}
2021-05-27T13:16:47.269+0800	INFO	Bootstrapping promEntry.	{"domain": "rk-demo", "entryName": "greeter-prom", "entryType": "GinPromEntry", "entryDescription": "Internal RK entry which implements prometheus client with Gin framework.", "path": "/metrics", "port": 8080}
2021-05-27T13:16:47.269+0800	INFO	Bootstrapping CommonServiceEntry.	{"domain": "rk-demo", "entryName": "greeter-commonService", "entryType": "GinCommonServiceEntry"}
2021-05-27T13:16:47.271+0800	INFO	Bootstrapping tvEntry.	{"domain": "rk-demo", "entryName": "greeter-tv", "entryType": "GinTvEntry", "path": "/rk/v1/tv/*item"}
2021-05-27T13:16:47.271+0800	INFO	Bootstrapping GinEntry.	{"domain": "rk-demo", "entryName": "greeter", "entryType": "GinEntry", "port": 8080, "interceptorsCount": 5, "swEnabled": true, "tlsEnabled": false, "commonServiceEnabled": true, "tvEnabled": true, "swPath": "/sw/", "promPath": "/metrics", "promPort": 8080}
2021-05-27T13:16:47.269+0800	INFO	Bootstrapping SwEntry.	{"domain": "rk-demo", "entryName": "greeter-sw", "entryType": "GinSwEntry", "jsonPath": "docs", "path": "/sw/", "port": 8080}
2021-05-27T13:16:47.269+0800	INFO	Bootstrapping promEntry.	{"domain": "rk-demo", "entryName": "greeter-prom", "entryType": "GinPromEntry", "entryDescription": "Internal RK entry which implements prometheus client with Gin framework.", "path": "/metrics", "port": 8080}
2021-05-27T13:16:47.269+0800	INFO	Bootstrapping CommonServiceEntry.	{"domain": "rk-demo", "entryName": "greeter-commonService", "entryType": "GinCommonServiceEntry"}
2021-05-27T13:16:47.271+0800	INFO	Bootstrapping tvEntry.	{"domain": "rk-demo", "entryName": "greeter-tv", "entryType": "GinTvEntry", "path": "/rk/v1/tv/*item"}
2021-05-27T13:16:47.271+0800	INFO	Bootstrapping GinEntry.	{"domain": "rk-demo", "entryName": "greeter", "entryType": "GinEntry", "port": 8080, "interceptorsCount": 5, "swEnabled": true, "tlsEnabled": false, "commonServiceEnabled": true, "tvEnabled": true, "swPath": "/sw/", "promPath": "/metrics", "promPort": 8080}
```

### Access Zap Logger from code
ZapLoggerEntry could be retrieved from rkentry.GlobalAppCtx.

```go
// *********** Access ZapLoggerEntry ***********
loggerEntry := rkentry.GlobalAppCtx.GetZapLoggerEntry("logger")
loggerEntry.GetLogger().Info("this is logger")
```

```text
2021-05-27T19:06:43.148+0800	INFO	this is logger
```

### Directory layout

```shell script
├── Makefile
├── README.md
├── boot.yaml
├── docs
|   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── logs
│   └── greeter.log
└── main.go
```

### boot.yaml
We are using the simplest way of boot.yaml with zap logger detail.
Available swagger configurations listed bellow.

ZapLoggerEntry follows zap and lumberjack YAML hierarchy, please refer to [zap](https://pkg.go.dev/go.uber.org/zap#section-documentation) and [lumberjack](https://github.com/natefinch/lumberjack) site for details.

| Name | Description | Default |
| ------ | ------ | ------ |
| zapLogger.name | Required. Name of zap logger entry | N/A |
| zapLogger.description | Description of zap logger entry. | N/A |
| zapLogger.zap.level | Level is the minimum enabled logging level. | info | 
| zapLogger.zap.development | Development puts the logger in development mode, which changes the behavior of DPanicLevel and takes stacktraces more liberally. | true |
| zapLogger.zap.disableCaller | DisableCaller stops annotating logs with the calling function's file name and line number. | false |
| zapLogger.zap.disableStacktrace | DisableStacktrace completely disables automatic stacktrace capturing. | true |
| zapLogger.zap.sampling | Sampling sets a sampling policy. | nil |
| zapLogger.zap.encoding | Encoding sets the logger's encoding. Valid values are "json" and "console", as well as any third-party encodings registered via RegisterEncoder. | console |
| zapLogger.zap.encoderConfig.messageKey | As name described. | msg |
| zapLogger.zap.encoderConfig.levelKey | As name described. | level |
| zapLogger.zap.encoderConfig.timeKey | As name described. | ts |
| zapLogger.zap.encoderConfig.nameKey | As name described. | logger |
| zapLogger.zap.encoderConfig.callerKey | As name described. | caller |
| zapLogger.zap.encoderConfig.functionKey | As name described. | "" |
| zapLogger.zap.encoderConfig.stacktraceKey | As name described. | stacktraceKey |
| zapLogger.zap.encoderConfig.lineEnding | As name described. | \n |
| zapLogger.zap.encoderConfig.timeEncoder | As name described. | iso8601 |
| zapLogger.zap.encoderConfig.durationEncoder | As name described. | string |
| zapLogger.zap.encoderConfig.callerEncoder | As name described. | "" |
| zapLogger.zap.encoderConfig.nameEncoder | As name described. | "" |
| zapLogger.zap.encoderConfig.consoleSeparator | As name described. | "" |
| zapLogger.zap.outputPaths | Output paths. | ["stdout"] |
| zapLogger.zap.errorOutputPaths | Output paths. | ["stderr"] |
| zapLogger.zap.initialFields | Output paths. | empty map |
| zapLogger.lumberjack.filename | Filename is the file to write logs to | It uses <processname>-lumberjack.log in os.TempDir() if empty. |
| zapLogger.lumberjack.maxsize | MaxSize is the maximum size in megabytes of the log file before it gets rotated. | 1024 |
| zapLogger.lumberjack.maxage | MaxAge is the maximum number of days to retain old log files based on the timestamp encoded in their filename. | 7 |
| zapLogger.lumberjack.maxbackups | axBackups is the maximum number of old log files to retain. | 3 |
| zapLogger.lumberjack.localtime | LocalTime determines if the time used for formatting the timestamps in backup files is the computer's local time. | true |
| zapLogger.lumberjack.compress | Compress determines if the rotated log files should be compressed using gzip. | true |

- Full zap logger config
```yaml
zapLogger:
  - name: zap-logger                      # Required
    description: "Description of entry"   # Optional
    zap:
      level: info                         # Optional, default: info, options: [debug, DEBUG, info, INFO, warn, WARN, dpanic, DPANIC, panic, PANIC, fatal, FATAL]
      development: true                   # Optional, default: true
      disableCaller: false                # Optional, default: false
      disableStacktrace: true             # Optional, default: true
      sampling:                           # Optional, default: empty map
        initial: 0
        thereafter: 0
      encoding: console                   # Optional, default: "console", options: [console, json]
      encoderConfig:
        messageKey: "msg"                 # Optional, default: "msg"
        levelKey: "level"                 # Optional, default: "level"
        timeKey: "ts"                     # Optional, default: "ts"
        nameKey: "logger"                 # Optional, default: "logger"
        callerKey: "caller"               # Optional, default: "caller"
        functionKey: ""                   # Optional, default: ""
        stacktraceKey: "msg"              # Optional, default: "msg"
        lineEnding: "\n"                  # Optional, default: "\n"
        levelEncoder: "capitalColor"      # Optional, default: "capitalColor", options: [capital, capitalColor, color, lowercase]
        timeEncoder: "iso8601"            # Optional, default: "iso8601", options: [rfc3339nano, RFC3339Nano, rfc3339, RFC3339, iso8601, ISO8601, millis, nanos]
        durationEncoder: "string"         # Optional, default: "string", options: [string, nanos, ms]
        callerEncoder: ""                 # Optional, default: ""
        nameEncoder: ""                   # Optional, default: ""
        consoleSeparator: ""              # Optional, default: ""
      outputPaths: [ "stdout" ]           # Optional, default: ["stdout"], stdout would be replaced if specified
      errorOutputPaths: [ "stderr" ]      # Optional, default: ["stderr"], stderr would be replaced if specified
      initialFields:                      # Optional, default: empty map
        key: "value"
    lumberjack:                           # Optional
      filename: "rkapp-event.log"         # Optional, default: It uses <processname>-lumberjack.log in os.TempDir() if empty.
      maxsize: 1024                       # Optional, default: 1024 (MB)
      maxage: 7                           # Optional, default: 7 (days)
      maxbackups: 3                       # Optional, default: 3 (days)
      localtime: true                     # Optional, default: true
      compress: true                      # Optional, default: true
```

- The simplest zap logger config
```yaml
---
rk:
  appName: rk-example-entry
  version: v0.0.1
  description: "this is description"
  keywords: ["rk", "golang"]
  homeUrl: "http://example.com"
  iconUrl: "http://example.com"
  docsUrl: ["http://example.com"]
  maintainers: ["rk-dev"]
zapLogger:
  - name: logger                      # Required
    zap:
      outputPaths: ["logs/greeter.log"]
gin:
  - name: greeter                     # Required
    port: 8080                        # Required
    description: "greeter server"
    sw:
      enabled: true
      jsonPath: "docs"
    prom:
      enabled: true
    tv:
      enabled: true
    commonService:
      enabled: true
    interceptors:
      loggingZap:
        enabled: true
      metricsProm:
        enabled: true
      basicAuth:
        enabled: true
        credentials:
         - "rk-user:rk-pass"
    logger:                                   # Optional
      zapLogger:                              # Optional
        ref: logger                           # Optional, default: logger of STDOUT, reference of logger entry declared above
```
