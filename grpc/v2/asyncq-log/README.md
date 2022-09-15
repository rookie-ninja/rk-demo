# Example
In this example, we will inject traceId from grpc context into asyncq server.

```shell
go get github.com/rookie-ninja/rk-boot/v2
go get github.com/rookie-ninja/rk-grpc/v2
```

## Quick start
### 3.Create server.yaml
- server.yaml

```yaml
grpc:
  - name: greeter
    enabled: true
    port: 8080
    middleware:
      trace:
        enabled: true
      logging:
        enabled: true
```


### 4.Create server.go and worker.go
Please refer to [grpc-server.go](grpc-server.go) and [worker.go](worker.go) 

### 5.Start redis, server and worker

```shell
$ docker run --name my-redis -p 6379:6379 -d redis
$ go run grpc-server.go
$ go run worker.go
```

### 4.Validation
#### 4.1 Send request to grpc server
grpc server will enqueue demo task

```shell
curl localhost:8080/v1/enqueue
{}
```

#### 4.2 Validate log worker and grpc-server
Two servers will have same traceId

- log from grpc-server

```shell
2022-09-16T04:41:46.761+0800    INFO    asyncq-log/grpc-server.go:54    enqueued task   {"traceId": "83928cd9869a15ef335f081fb539d5d6", "id": "81aa8ed8-3244-477e-9ddc-3eada669216d", "queue": "default"}
------------------------------------------------------------------------
ids={"eventId":"c6880511-f229-4bf1-a234-29976eccea7c","traceId":"83928cd9869a15ef335f081fb539d5d6"}
...
```

- log from worker

```shell
2022-09-16T04:41:47.057+0800    INFO    task/task.go:36 handle demo task        {"traceId": "83928cd9869a15ef335f081fb539d5d6"}
```

