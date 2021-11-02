# Basic tutorial
In this example, we will show all YAML config options by functionality block.

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [Application metadata](#application-metadata)
- [Zap logger](#zap-logger)
- [Event logger](#event-logger)
- [Config](#config)
- [Credential](#credential)
- [Certificates](#certificates)
- [Echo](#echo)
  - [Echo server](#echo-server)
  - [Swagger](#swagger)
  - [Common service](#common-service)
  - [TV](#tv)
  - [Prometheus client](#prometheus-client)
  - [Middleware/Interceptor](#middlewareinterceptor)
    - [logging](#logging)
    - [Metrics prometheus](#metrics-prometheus)
    - [Auth](#auth)
    - [Meta](#meta)
    - [Tracing](#tracing)
    - [Rate limit](#rate-limit)
    - [Timeout](#timeout)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Application metadata
This will be in the response of /rk/v1/info which is accessible if commonService was enabled.

```yaml
---
app:
  description: "this is description"                      # Optional, default: ""
  keywords: ["rk", "golang"]                              # Optional, default: []
  homeUrl: "http://example.com"                           # Optional, default: ""
  iconUrl: "http://example.com"                           # Optional, default: ""
  docsUrl: ["http://example.com"]                         # Optional, default: []
  maintainers: ["rk-dev"]                                 # Optional, default: []
```

- Access from code

```go
rkentry.GlobalAppCtx.GetAppInfoEntry()
```

## Zap logger
zapLogger.zap and zapLogger.lumberjack is compatible with the official configuration.

- [zap](https://github.com/uber-go/zap/blob/master/config.go)
- [lumberjack](https://github.com/natefinch/lumberjack/blob/v2.0/lumberjack.go)

```yaml
zapLogger:
  - name: zap-logger                                      # Required
    description: "Description of entry"                   # Optional
    zap:
      level: info                                         # Optional, default: info, options: [debug, DEBUG, info, INFO, warn, WARN, dpanic, DPANIC, panic, PANIC, fatal, FATAL]
      development: true                                   # Optional, default: true
      disableCaller: false                                # Optional, default: false
      disableStacktrace: true                             # Optional, default: true
      sampling:
        initial: 0                                        # Optional, default: 0
        thereafter: 0                                     # Optional, default: 0
      encoding: console                                   # Optional, default: "console", options: [console, json]
      encoderConfig:
        messageKey: "msg"                                 # Optional, default: "msg"
        levelKey: "level"                                 # Optional, default: "level"
        timeKey: "ts"                                     # Optional, default: "ts"
        nameKey: "logger"                                 # Optional, default: "logger"
        callerKey: "caller"                               # Optional, default: "caller"
        functionKey: ""                                   # Optional, default: ""
        stacktraceKey: "msg"                              # Optional, default: "msg"
        lineEnding: "\n"                                  # Optional, default: "\n"
        levelEncoder: "capitalColor"                      # Optional, default: "capitalColor", options: [capital, capitalColor, color, lowercase]
        timeEncoder: "iso8601"                            # Optional, default: "iso8601", options: [rfc3339nano, RFC3339Nano, rfc3339, RFC3339, iso8601, ISO8601, millis, nanos]
        durationEncoder: "string"                         # Optional, default: "string", options: [string, nanos, ms]
        callerEncoder: ""                                 # Optional, default: ""
        nameEncoder: ""                                   # Optional, default: ""
        consoleSeparator: ""                              # Optional, default: ""
      outputPaths: [ "stdout" ]                           # Optional, default: ["stdout"], stdout would be replaced if specified
      errorOutputPaths: [ "stderr" ]                      # Optional, default: ["stderr"], stderr would be replaced if specified
      initialFields:                                      # Optional, default: empty map
        key: "value"
    lumberjack:
      filename: "rkapp.log"                               # Optional, default: It uses <processname>-lumberjack.log in os.TempDir() if empty.
      maxsize: 1024                                       # Optional, default: 1024 (MB)
      maxage: 7                                           # Optional, default: 7 (days)
      maxbackups: 3                                       # Optional, default: 3 (days)
      localtime: true                                     # Optional, default: true
      compress: true                                      # Optional, default: true
```

- Access from code

```go
rkentry.GlobalAppCtx.GetZapLoggerEntry("zap-logger")
```

## Event logger

```yaml
eventLogger:
  - name: event-logger                                    # Required
    description: "Description of entry"                   # Optional
    encoding: "json"                                      # Optional, default: console, options: [json, console]
    outputPaths: []                                       # Optional, default: ["stdout"], stdout would be replaced if specified
    lumberjack:
      filename: "rkapp.log"                               # Optional, default: It uses <processname>-lumberjack.log in os.TempDir() if empty.
      maxsize: 1024                                       # Optional, default: 1024 (MB)
      maxage: 7                                           # Optional, default: 7 (days)
      maxbackups: 3                                       # Optional, default: 3 (days)
      localtime: true                                     # Optional, default: true
      compress: true                                      # Optional, default: true
```

- Access from code

```go
rkentry.GlobalAppCtx.GetEventLoggerEntry("event-logger")
```

## Config
Provide configuration file path which will be parsed by [viper](https://github.com/spf13/viper).

Supported file type:
- JSON
- TOML
- YAML
- HCL
- envfile
- Java properties config files

```yaml
config:
  - name: my-config                                       # Required
    locale: "*::*::*::*"                                  # Required, default: ""
    path: config/default.yaml                             # Required
    description: "Description of entry"                   # Optional
```

- Access from code

```go
rkentry.GlobalAppCtx.GetConfigEntry("my-config")
```

## Credential
Retrieve credential from local or remote service.

- localFs
- remoteFs
- etcd
- consul

```yaml
cred:
  - name: "local-cred"                                    # Required
    description: "Description of entry"                   # Optional
    provider: "localFs"                                   # Required, etcd, consul, localFs, remoteFs are supported options
    locale: "*::*::*::*"                                  # Required, default: ""
    paths:                                                # Optional
      - "example/boot/full/cred.yaml"
```

- Access from code

```go
rkentry.GlobalAppCtx.GetCredEntry("local-cred")
```

## Certificates
Retrieve certificates from local or remote service.

- localFs
- remoteFs
- etcd
- consul

```yaml
cert:
  - name: "local-cert"                                    # Required
    description: "Description of entry"                   # Optional
    provider: "localFs"                                   # Required, etcd, consul, localFs, remoteFs are supported options
    locale: "*::*::*::*"                                  # Required, default: ""
    serverCertPath: "cert/server.pem"                     # Optional, default: "", path of certificate on local FS
    serverKeyPath: "cert/server-key.pem"                  # Optional, default: "", path of certificate on local FS
```

- Access from code

```go
rkentry.GlobalAppCtx.GetCertEntry("local-cert")
```

## Echo

### Echo server
Echo server general configuration.

```yaml
echo:
  - name: greeter                                         # Required
    port: 8080                                            # Required
    enabled: true                                         # Required
    description: "greeter server"                         # Optional, default: ""
    cert:
      ref: "local-cert"                                   # Optional, default: "", reference of cert entry declared above
    logger:
      zapLogger:
        ref: zap-logger                                   # Optional, default: logger of STDOUT, reference of logger entry declared above
      eventLogger:
        ref: event-logger                                 # Optional, default: logger of STDOUT, reference of logger entry declared above
```

- Access from code

```go
	// Create a new boot instance.
	boot := rkboot.NewBoot()

	// Register handler
	boot.GetEchoEntry("greeter").Echo.GET("/v1/greeter", Greeter)
```

### Swagger
**path** is web url.

**jsonPath** is local file path where swagger configuration exists.

**headers** are the headers that will be sent to user in the response.

```yaml
echo:
  - name: greeter                                          # Required
    port: 8080                                             # Required
    enabled: true                                          # Required
    sw:
      enabled: true                                        # Optional, default: false
      path: "sw"                                           # Optional, default: "sw"
      jsonPath: ""                                         # Optional, default: ""
      headers: ["sw:rk"]                                   # Optional, default: []
```

- Access from code

```go
boot.GetEchoEntry("greeter").SwEntry
```

### Common service

```yaml
echo:
  - name: greeter                                          # Required
    port: 8080                                             # Required
    enabled: true                                          # Required
    commonService:
      enabled: true                                        # Optional, default: false
```

- Access from code

```go
boot.GetEchoEntry("greeter").CommonServiceEntry
```

### TV

```yaml
echo:
  - name: greeter                                          # Required
    port: 8080                                             # Required
    enabled: true                                          # Required
    tv:
      enabled: true                                        # Optional, default: false
```

- Access from code

```go
boot.GetEchoEntry("greeter").TvEntry
```

### Prometheus client

```yaml
echo:
  - name: greeter                                          # Required
    port: 8080                                             # Required
    enabled: true                                          # Required
    prom:
      enabled: true                                        # Optional, default: false
      path: ""                                             # Optional, default: "metrics"
      pusher:                                              # pushgateway configuration
        enabled: false                                     # Optional, default: false
        jobName: "greeter-pusher"                          # Required
        remoteAddress: "localhost:9091"                    # Required
        basicAuth: "user:pass"                             # Optional, default: ""
        intervalMs: 10000                                  # Optional, default: 1000
        cert:                                              # Optional
          ref: "local-test"                                # Optional, default: "", reference of cert entry declared above
```

- Access from code

```go
boot.GetEchoEntry("greeter").PromEntry
```

### Middleware/Interceptor

#### logging
Logging middleware will use log instance from **echo.logger.zapLogger.ref** and **echo.logger.eventLogger.ref**.

User can override encoding and output paths.

```yaml
echo:
  - name: greeter                                          # Required
    port: 8080                                             # Required
    enabled: true                                          # Required
    interceptors:
      loggingZap:
        enabled: true                                      # Optional, default: false
        zapLoggerEncoding: "json"                          # Optional, default: "console"
        zapLoggerOutputPaths: ["logs/app.log"]             # Optional, default: ["stdout"]
        eventLoggerEncoding: "json"                        # Optional, default: "console"
        eventLoggerOutputPaths: ["logs/event.log"]         # Optional, default: ["stdout"]
```

#### Metrics prometheus
**echo.prom** must be enabled in order to use it.

```yaml
echo:
  - name: greeter                                          # Required
    port: 8080                                             # Required
    enabled: true                                          # Required
    interceptors:
      metricsProm:
        enabled: true                                      # Optional, default: false

```

#### Auth

```yaml
echo:
  - name: greeter                                          # Required
    port: 8080                                             # Required
    enabled: true                                          # Required
    interceptors:
      auth:
        enabled: true                                      # Optional, default: false
        basic:
          - "user:pass"                                    # Optional, default: []
        ignorePrefix:
          - "/rk/v1"                                       # Optional, default: []
        apiKey:
          - "keys"                                         # Optional, default: []
```

#### Meta

```yaml
echo:
  - name: greeter                                          # Required
    port: 8080                                             # Required
    enabled: true                                          # Required
    interceptors:
      meta:
        enabled: true                                      # Optional, default: false
        prefix: "rk"                                       # Optional, default: "rk"
```

#### Tracing

```yaml
echo:
  - name: greeter                                          # Required
    port: 8080                                             # Required
    enabled: true                                          # Required
    interceptors:
      tracingTelemetry:
        enabled: true                                      # Optional, default: false
        exporter:                                          # Optional, default will create a stdout exporter
          file:
            enabled: true                                  # Optional, default: false
            outputPath: "logs/trace.log"                   # Optional, default: stdout
          jaeger:
            agent:
              enabled: false                               # Optional, default: false
              host: ""                                     # Optional, default: localhost
              port: 0                                      # Optional, default: 6831
            collector:
              enabled: true                                # Optional, default: false
              endpoint: ""                                 # Optional, default: http://localhost:14268/api/traces
              username: ""                                 # Optional, default: ""
              password: ""                                 # Optional, default: ""
```

#### Rate limit

```yaml
echo:
  - name: greeter                                          # Required
    port: 8080                                             # Required
    enabled: true                                          # Required
    interceptors:
      rateLimit:
        enabled: false                                     # Optional, default: false
        algorithm: "leakyBucket"                           # Optional, default: "tokenBucket", options: [tokenBucket, leakyBucket]
        reqPerSec: 100                                     # Optional, default: 1000000
        paths:
          - path: "/rk/v1/healthy"                         # Optional, default: ""
            reqPerSec: 0                                   # Optional, default: 1000000
```

#### Timeout

```yaml
echo:
  - name: greeter                                          # Required
    port: 8080                                             # Required
    enabled: true                                          # Required
    interceptors:
      timeout:
        enabled: false                                     # Optional, default: false
        timeoutMs: 5000                                    # Optional, default: 5000
        paths:
          - path: "/rk/v1/healthy"                         # Optional, default: ""
            timeoutMs: 1000                                # Optional, default: 5000
```