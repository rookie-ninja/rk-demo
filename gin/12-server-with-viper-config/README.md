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
- Viper Config

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

### Access config form code
ConfigEntry with name already registered into rkentry-GlobalApCtx.

```go
// ************* Access ConfigEntry *************
// Access config entry
configEntry := rkentry.GlobalAppCtx.GetConfigEntry("my-config")
// Get config value from viper instance
fmt.Println(cast.ToString(configEntry.GetViper().Get("rk-key")))
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
We are using the simplest way of boot.yaml with config detail.
Available swagger configurations listed bellow.

| Name | Description | Required | Default |
| ------ | ------ | ------ | ------ |
| config.name | Name of config entry. | required | config-<random string> |
| config.path | File path of config file, could be either relative or absolute path | required | "" | 
| config.locale | <realm>::<region>::<az>::<domain> | required | "" |
| config.description | Description of config entry. | "" |

- Full event logger config
```yaml
---
config:
  - name: my-config                       # Required
    locale: "*::*::*::*"                  # Required
    path: example/my-config.yaml          # Required
    description: "Description of entry"   # Optional
```

- The simplest config for configEntry
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
config:
  - name: my-config
    path: configs/my-config.yaml
    locale: "*::*::*::*"
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
