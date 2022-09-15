# Example
In this example, we will start two gRPC servers serverA and serverB and implement distributed logging.

```shell
go get github.com/rookie-ninja/rk-boot/v2
go get github.com/rookie-ninja/rk-grpc/v2
```

## Quick start
### 3.Create bootA.yaml and bootB.yaml
- bootA.yaml

```yaml
grpc:
  - name: serverA
    enabled: true
    port: 1949
    middleware:
      trace:
        enabled: true
      logging:
        enabled: true
```

- bootB.yaml

```yaml
grpc:
  - name: serverB
    enabled: true
    port: 2008
    middleware:
      trace:
        enabled: true
      logging:
        enabled: true
```

### 4.Create serverA.go and serverB.go
Please refer to [serverA.go](serverA.go) and [serverB.go](serverB.go) 

### 5.Start serverA and serverB

```shell
$ go run serverA.go
$ go run serverB.go
```

### 4.Validation
#### 4.1 Send request to serverA
Since grpc-gateway is automatically enabled by default, we will use curl to send request.

```shell
curl "localhost:1949/v1/hello?name=rk-dev"
{"message":"Hello rk-dev!"}
```

#### 4.2 Validate log from serverA and serverB
Two servers will have same traceId

- log from ServerA

```shell
------------------------------------------------------------------------
ids={"eventId":"f5204ea9-b727-4c93-9b22-beea6fae147c","traceId":"e04a5378e5216becf8afb947a1428a7e"}
...
```

- log from ServerB

```shell
------------------------------------------------------------------------
ids={"eventId":"0dab9066-0352-4f1a-8221-356bc4e661bc","traceId":"e04a5378e5216becf8afb947a1428a7e"}
....
```