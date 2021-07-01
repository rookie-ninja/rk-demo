# Gin server with rk-boot
This is an example demonstrate how to configure boot.yaml file to start a gin server with bellow functionality.

| Name | Description |
| ---- | ---- |
| Start with YAML | Start service with YAML config. |
| Start with code | Start service from code. |
| Gin Service | Gin service. |
| Swagger Service | Swagger UI. |
| Common Service | List of common API available on Gin. |
| TV Service | A Web UI shows application and environment information. |
| Metrics interceptor | Collect RPC metrics and export as prometheus client. |
| Log interceptor | Log every RPC requests as event with rk-query. |
| Trace interceptor | Collect RPC trace and export it to stdout, file or jaeger. |
| Panic interceptor | Recover from panic for RPC requests and log it. |
| Meta interceptor | Send application metadata as header to client. |
| Auth interceptor | Support [Basic Auth], [Bearer Token] and [API Key] authorization types. |

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [Quick start](#quick-start)
  - [boot.yaml](#bootyaml)
  - [Start server](#start-server)
  - [Log output](#log-output)
  - [Directory layout](#directory-layout)
- [Enable common service](#enable-common-service)
  - [boot.yaml](#bootyaml-1)
  - [Send request](#send-request)
- [Enable swagger UI](#enable-swagger-ui)
  - [boot.yaml](#bootyaml-2)
  - [Access swagger UI](#access-swagger-ui)
- [Enable prometheus client](#enable-prometheus-client)
  - [boot.yaml](#bootyaml-3)
  - [Access prometheus client](#access-prometheus-client)
- [Enable TV](#enable-tv)
  - [boot.yaml](#bootyaml-4)
  - [Access TV](#access-tv)
- [Enable log interceptor(middleware)](#enable-log-interceptormiddleware)
  - [boot.yaml](#bootyaml-5)
  - [Send request](#send-request-1)
  - [Log output](#log-output-1)
- [Enable metrics interceptor(middleware)](#enable-metrics-interceptormiddleware)
  - [boot.yaml](#bootyaml-6)
  - [Send request](#send-request-2)
  - [Output](#output)
- [Enable meta interceptor(middleware)](#enable-meta-interceptormiddleware)
  - [boot.yaml](#bootyaml-7)
  - [Send request](#send-request-3)
- [Enable panic interceptor(middleware)](#enable-panic-interceptormiddleware)
- [Enable auth interceptor(middleware)](#enable-auth-interceptormiddleware)
  - [boot.yaml](#bootyaml-8)
  - [Send request](#send-request-4)
- [Enable trace interceptor(middleware)](#enable-trace-interceptormiddleware)
  - [boot.yaml](#bootyaml-9)
  - [Start jaeger all-in-one locally](#start-jaeger-all-in-one-locally)
  - [Send request](#send-request-5)
  - [Validate tracing](#validate-tracing)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Quick start
### boot.yaml
We are using boot.yaml which contains only name and port to start gin server.

| name | description | type | default value |
| ------ | ------ | ------ | ------ |
| gin.name | Name of gin server entry | string | N/A |
| gin.port | Port of server | integer | nil, server won't start |

```yaml
---
gin:
  - name: greeter                     # Required
    port: 8080                        # Required
```

### Start server
Run main.go in the terminal or run it from your IDE directly.

```go
$ go run main.go 
```

### Log output
It will print logs something like bellow:
- The first line printed from default [ZapLoggerEntry](https://github.com/rookie-ninja/rk-entry#zaploggerentry) which is commonly used logger format.
- The rest lines of logs are human readable format of [EventLoggerEntry](https://github.com/rookie-ninja/rk-entry#eventloggerentry) which is used to log every event.

```text
2021-07-01T05:44:11.043+0800    INFO    boot/gin_entry.go:630   Bootstrapping GinEntry. {"eventId": "ad44aa09-c8fc-4417-b91f-edc2434bc8ea", "entryName": "greeter", "entryType": "GinEntry", "port": 8080, "interceptorsCount": 1, "swEnabled": false, "tlsEnabled": false, "commonServiceEnabled": false, "tvEnabled": false}
------------------------------------------------------------------------
endTime=2021-07-01T05:44:11.043854+08:00
startTime=2021-07-01T05:44:11.043764+08:00
elapsedNano=89638
timezone=CST
ids={"eventId":"ad44aa09-c8fc-4417-b91f-edc2434bc8ea"}
app={"appName":"rk-demo","appVersion":"master-b233d43","entryName":"greeter","entryType":"GinEntry"}
env={"arch":"amd64","az":"*","domain":"*","hostname":"lark.local","localIP":"10.8.0.2","os":"darwin","realm":"*","region":"*"}
payloads={"commonServiceEnabled":false,"entryName":"greeter","entryType":"GinEntry","interceptorsCount":1,"port":8080,"swEnabled":false,"tlsEnabled":false,"tvEnabled":false}
error={}
counters={}
pairs={}
timing={}
remoteAddr=localhost
operation=bootstrap
resCode=OK
eventStatus=Ended
EOE
```

### Directory layout
boot.yaml is the bootstrap config file for rk-boot, rk-boot will read this file to start Gin server.
We locate boot.yaml file in the root working directory. As a result, we didn't specify file path of bootstrapper config file 
in the main.go function. Because rk-boot will look for bootstrapper file in the root working directory named as boot.yaml
if not specified.

```shell script
├── boot.yaml
└── main.go
```



## Enable common service
[Common service](https://github.com/rookie-ninja/rk-gin#common-service) is collection of built in utility APIs.

### boot.yaml
Available options.

| name | description | type | default value |
| ------ | ------ | ------ | ------ |
| gin.commonService.enabled | Enable embedded common service | boolean | false |

```yaml
---
gin:
  - name: greeter
    port: 8080
    # Enable common service
    commonService:
      enabled: true
```

### Send request
Pick one API described bellow.

```shell script
$ curl localhost:8080/rk/v1/healthy
{"healthy":true}
```

| Path | Description |
| ---- | ---- |
| /rk/v1/apis | /rk/v1/apis |
| /rk/v1/certs | List CertEntry |
| /rk/v1/configs | List ConfigEntry |
| /rk/v1/deps | List dependencies related application |
| /rk/v1/entries | List all Entry |
| /rk/v1/gc | Trigger Gc |
| /rk/v1/healthy | Get application healthy status |
| /rk/v1/info | Get application and process info |
| /rk/v1/license | Get license related application |
| /rk/v1/logs | List logger related entries |
| /rk/v1/git | Get git information |
| /rk/v1/readme | Get README file |
| /rk/v1/req | List prometheus metrics of requests |
| /rk/v1/sys | Get OS stat |
| /rk/v1/tv | Get HTML page of /tv |

## Enable swagger UI
[Swagger](https://swagger.io/tools/swagger-ui/) will be accessible as soon as common service enabled since common service will
register swagger UI automatically.

### boot.yaml
Available options.

| name | description | type | default value |
| ------ | ------ | ------ | ------ |
| gin.sw.enabled | Enable swagger service over gin server | boolean | false |
| gin.sw.path | The path access swagger service from web | string | /sw |
| gin.sw.jsonPath | Where the swagger.json files are stored locally | string | "" |
| gin.sw.headers | Headers would be sent to caller as scheme of [key:value] | []string | [] |

```yaml
---
gin:
  - name: greeter
    port: 8080
    # Enable common service
    commonService:
      enabled: true
    # Enable swagger UI
    sw:
      enabled: true
```

### Access swagger UI
- http://localhost:8080/sw/

![swagger UI](img/enable-swagger.png)

## Enable prometheus client
Enable [prometheus client](https://prometheus.io/docs/instrumenting/writing_clientlibs/) from boot.yaml where user would access local prometheus metrics.

### boot.yaml
Available options.

| name | description | type | default value |
| ------ | ------ | ------ | ------ |
| gin.prom.enabled | Enable prometheus | boolean | false |
| gin.prom.path | Path of prometheus | string | /metrics |
| gin.prom.pusher.enabled | Enable prometheus pusher | bool | false |
| gin.prom.pusher.jobName | Job name would be attached as label while pushing to remote [pushgateway](https://github.com/prometheus/pushgateway) | string | "" |
| gin.prom.pusher.remoteAddress | PushGateWay address, could be form of http://x.x.x.x or x.x.x.x | string | "" |
| gin.prom.pusher.intervalMs | Push interval in milliseconds | string | 1000 |
| gin.prom.pusher.basicAuth | Basic auth used to interact with remote [pushgateway](https://github.com/prometheus/pushgateway), form of [user:pass] | string | "" |
| gin.prom.pusher.cert.ref | Reference of [CertEntry](https://github.com/rookie-ninja/rk-entry#certentry) | string | "" |

```yaml
---
gin:
  - name: greeter
    port: 8080
    # Enable prometheus client
    prom:
      enabled: true
```

### Access prometheus client
- http://localhost:8080/metrics

![prom client](img/enable-prom-client.png)

## Enable TV
RK TV is a built in web UI which shows application information. All of js, images, css files were builtin assets, as a result, user don't need internet access.

### boot.yaml
Available options.

| name | description | type | default value |
| ------ | ------ | ------ | ------ |
| gin.tv.enabled | Enable RK TV | boolean | false |

```yaml
---
gin:
  - name: greeter
    port: 8080
    # Enable TV
    tv:
      enabled : true
```

### Access TV
- http://localhost:8080/rk/v1/tv/

![tv](img/enable-tv.png)

## Enable log interceptor(middleware)
Log interceptor(middleware) will intercept each RPC and log request metadata into [eventLogger](https://github.com/rookie-ninja/rk-gin#log).

### boot.yaml
Available options.

| name | description | type | default value |
| ------ | ------ | ------ | ------ |
| gin.interceptors.loggingZap.enabled | Enable log interceptor | boolean | false |

```yaml
---
gin:
  - name: greeter
    port: 8080
    # Enable common service in order to send request for testing logging interceptor(middleware)
    commonService:
      enabled: true
    interceptors:
      # Enable logging interceptor(middleware)
      loggingZap:
        enabled: true
```

### Send request
```shell script
$ curl -X GET "http://localhost:8080/rk/v1/healthy"
{"healthy":true}
```

### Log output
```shell script
------------------------------------------------------------------------
endTime=2021-07-01T17:04:32.91871+08:00
startTime=2021-07-01T17:04:32.918648+08:00
elapsedNano=61990
timezone=CST
ids={"eventId":"4fe0d8e0-8086-416b-a45e-b2350647b832"}
app={"appName":"rk-demo","appVersion":"master-b233d43","entryName":"greeter","entryType":"GinEntry"}
env={"arch":"amd64","az":"*","domain":"*","hostname":"lark.local","localIP":"10.8.0.2","os":"darwin","realm":"*","region":"*"}
payloads={"apiMethod":"GET","apiPath":"/rk/v1/healthy","apiProtocol":"HTTP/1.1","apiQuery":"","userAgent":"curl/7.64.1"}
error={}
counters={}
pairs={}
timing={}
remoteAddr=localhost:57798
operation=/rk/v1/healthy
resCode=200
eventStatus=Ended
EOE
```

## Enable metrics interceptor(middleware)
Metrics interceptor(middleware) will intercept each RPC and record request ResCode, Error and ElapsedTime into prometheus client.

_In order record metrics, we need to enable prometheus client either._

### boot.yaml
Available options.

| name | description | type | default value |
| ------ | ------ | ------ | ------ |
| gin.interceptors.metricsProm.enabled | Enable metrics interceptor | boolean | false |

```yaml
---
gin:
  - name: greeter
    port: 8080
    # Enable common service in order to send request for testing metrics interceptor(middleware)
    commonService:
      enabled: true
    # Enable prom client in order to monitoring prometheus client
    prom:
      enabled: true
    interceptors:
      metricsProm:
        enabled: true
```

### Send request
```shell script
$ curl -X GET "http://localhost:8080/rk/v1/healthy"
{"healthy":true}
```

### Output
- http://localhost:8080/metrics

![prom](img/enable-prom-inter.png)

## Enable meta interceptor(middleware)
Meta interceptor(middleware) will intercept each RPC and send application metadata as header to client with response.

| Header key | Description |
| ---- | ---- |
| X-Request-Id | Request id generated by the interceptor. |
| X-[Prefix]-App | Application name. |
| X-[Prefix]-App-Version | Version of application. |
| X-[Prefix]-App-Unix-Time | Unix time of current application. |
| X-[Prefix]-Request-Received-Time | Time of current request received by application. |

### boot.yaml
```yaml
---
gin:
  - name: greeter
    port: 8080
    # Enable common service in order to send request for testing meta interceptor(middleware)
    commonService:
      enabled: true
    interceptors:
      # Enable meta interceptor(middleware)
      meta:
        enabled: true
```

### Send request
```shell script
$ curl -vs -X GET "http://localhost:8080/rk/v1/healthy"
...
< X-Request-Id: e3dfd422-15c8-4e56-8893-61be5b88f64f
< X-Rk-App-Name: rk-demo
< X-Rk-App-Unix-Time: 2021-07-01T17:32:12.372285+08:00
< X-Rk-App-Version: master-b233d43
< X-Rk-Received-Time: 2021-07-01T17:32:12.372285+08:00
...
{"healthy":true}
```

## Enable panic interceptor(middleware)
By default, [panic](https://github.com/rookie-ninja/rk-gin/tree/master/example/interceptor/panic) will enabled with rk-boot.

## Enable auth interceptor(middleware)
Auth interceptor will validate bellow authorizations.

| Type | Description | Example |
| ---- | ---- | ---- |
| Basic Auth | The client sends HTTP requests with the Authorization header that contains the word Basic, followed by a space and a base64-encoded(non-encrypted) string username: password. | Authorization: Basic AXVubzpwQDU1dzByYM== |
| Bearer Token | Commonly known as token authentication. It is an HTTP authentication scheme that involves security tokens called bearer tokens. | Authorization: Bearer [token] |
| API Key | An API key is a token that a client provides when making API calls. With API key auth, you send a key-value pair to the API in the request headers. | X-API-Key: abcdefgh123456789 | 

### boot.yaml
Available options.

| name | description | type | default value |
| ------ | ------ | ------ | ------ |
| gin.interceptors.auth.enabled | Enable auth interceptor | boolean | false |
| gin.interceptors.auth.basic | Basic auth credentials as scheme of <user:pass> | []string | [] |
| gin.interceptors.auth.bearer | Bearer auth tokens | []string | [] |
| gin.interceptors.auth.api | API key | []string | [] |

```yaml
---
gin:
  - name: greeter
    port: 8080
    interceptors:
      # Enable basic auth interceptor(middleware)
      auth:
        enabled: true
        basic:
          - "user:pass"
```

### Send request
By default, common services won't validate auth. Please add your own API which will authorised by default.

- Example
```shell script
$ curl "localhost:8080/rk/example/greeter?name=rk-dev"
# Pretty print manually.
{
    "error":{
        "code":401,
        "status":"Unauthorized",
        "message":"Missing authorization",
        "details":[]
    }
}
```

## Enable trace interceptor(middleware)
Trace interceptor will intercept each RPC and collect API metadata and export them to stdout, file or jaeger collector.

_The [OpenTelemetry](https://opentelemetry.io/) for jaeger agent export isn't stable yet, as a result, we are sending to jaeger collector directly._

### boot.yaml
Available options.

| name | description | type | default value |
| ------ | ------ | ------ | ------ |
| gin.interceptors.tracingTelemetry.enabled | Enable tracing interceptor | boolean | false |
| gin.interceptors.exporter.file.enabled | Enable file exporter | boolean | RK |
| gin.interceptors.exporter.file.outputPath | Export tracing info to files | string | stdout |
| gin.interceptors.exporter.jaeger.enabled | Export tracing info jaeger | boolean | false |
| gin.interceptors.exporter.jaeger.collectorEndpoint | As name described | string | localhost:16368/api/trace |
| gin.interceptors.exporter.jaeger.collectorUsername | As name described | string | "" |
| gin.interceptors.exporter.jaeger.collectorPassword | As name described | string | "" |

```yaml
---
gin:
  - name: greeter
    port: 8080
    # Enable common service in order to send request to server.
    commonService:
      enabled: true
    interceptors:
      tracingTelemetry:
        enabled: true
        exporter:
          # Export to local jaeger collector
          jaeger:
            enabled: true
```

### Start jaeger all-in-one locally
Please refer to jaeger official site: [Link](https://www.jaegertracing.io/docs/1.23/getting-started/)

```shell script
$ docker run -d --name jaeger \
  -e COLLECTOR_ZIPKIN_HOST_PORT=:9411 \
  -p 5775:5775/udp \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 14268:14268 \
  -p 14250:14250 \
  -p 9411:9411 \
  jaegertracing/all-in-one:1.23
```

### Send request
```shell script
$ curl -X GET "http://localhost:8080/rk/v1/healthy"
```

### Validate tracing
Open jaeger UI at http://localhost:16686

![jaeger](img/enable-jaeger-inter.png)