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
- Event Logger

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
New greeter-event.log file would be created automatically.

```text
------------------------------------------------------------------------
endTime=2021-05-27T14:57:04.412178+08:00
startTime=2021-05-27T14:57:04.411894+08:00
elapsedNano=283869
hostname=lark.local
timing={}
counter={}
pair={}
error={}
field={"entryName":"greeter-sw","entryType":"GinSwEntry","jsonPath":"docs","path":"/sw/","port":8080}
remoteAddr=lark.local
appName=rk-example-entry
appVersion=v0.0.1
entryName=greeter-sw
entryType=GinSwEntry
locale=unknown::unknown::unknown::unknown
operation=bootstrap
eventStatus=Ended
timezone=CST
os=darwin
arch=amd64
EOE
```

### Access Event Logger from code
EventLoggerEntry could be retrieved from rkentry.GlobalAppCtx.

```go
// *********** Access EventLoggerEntry ***********
eventLoggerEntry := rkentry.GlobalAppCtx.GetEventLoggerEntry("event-logger")
// EventHelper is a helper instance user can easily create and finish event.
helper := eventLoggerEntry.GetEventHelper()
// Start event
event := helper.Start("my-operation")
// Finish event
helper.Finish(event)
```

```text
------------------------------------------------------------------------
endTime=2021-05-27T19:10:37.760227+08:00
startTime=2021-05-27T19:10:37.760227+08:00
elapsedNano=127
hostname=lark.local
timing={}
counter={}
pair={}
error={}
field={}
remoteAddr=lark.local
appName=rk-example-entry
appVersion=v0.0.1
locale=unknown::unknown::unknown::unknown
operation=my-operation
eventStatus=Ended
timezone=CST
os=darwin
arch=amd64
EOE
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
│   └── greeter-event.log
└── main.go
```

### boot.yaml
We are using the simplest way of boot.yaml with event logger detail.
Available swagger configurations listed bellow.

| Name | Description | Default |
| ------ | ------ | ------ |
| rk.appName | Application name which refers to go process. | rkapp | 
| eventLogger.name | Required. Name of event logger entry. | N/A |
| eventLogger.description | Description of event logger entry. | N/A |
| eventLogger.format | Format of event logger, RK & JSON is supported. Please refer rkquery.RK & rkquery.JSON. | RK | 
| eventLogger.outputPaths | Output paths of event logger, stdout would be the default one if not provided. If one of output path was provided, then stdout would be omitted. Output path could be relative or absolute paths either. | stdout |
| eventLogger.lumberjack.filename | Filename is the file to write logs to | It uses <processname>-lumberjack.log in os.TempDir() if empty. |
| eventLogger.lumberjack.maxsize | MaxSize is the maximum size in megabytes of the log file before it gets rotated. | 1024 |
| eventLogger.lumberjack.maxage | MaxAge is the maximum number of days to retain old log files based on the timestamp encoded in their filename. | 7 |
| eventLogger.lumberjack.maxbackups | axBackups is the maximum number of old log files to retain. | 3 |
| eventLogger.lumberjack.localtime | LocalTime determines if the time used for formatting the timestamps in backup files is the computer's local time. | true |
| eventLogger.lumberjack.compress | Compress determines if the rotated log files should be compressed using gzip. | true |

- Full event logger config
```yaml
---
eventLogger:
  - name: event-logger                    # Required
    description: "Description of entry"   # Optional
    format: RK                            # Optional, default: RK, options: RK and JSON
    outputPaths: ["stdout"]               # Optional
    lumberjack:                           # Optional
      filename: "rkapp-event.log"         # Optional, default: It uses <processname>-lumberjack.log in os.TempDir() if empty.
      maxsize: 1024                       # Optional, default: 1024 (MB)
      maxage: 7                           # Optional, default: 7 (days)
      maxbackups: 3                       # Optional, default: 3 (days)
      localtime: true                     # Optional, default: true
      compress: true                      # Optional, default: true
```

- The simplest event logger config
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
  - name: logger                            # Required
    zap:
      outputPaths: ["logs/greeter.log"]
eventLogger:
  - name: event-logger                      # Required
    outputPaths: ["logs/greeter-event.log"]
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
      eventLogger:                            # Optional
        ref: event-logger                     # Optional, default: logger of STDOUT, reference of logger entry declared above
```
