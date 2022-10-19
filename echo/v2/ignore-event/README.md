# Ignoring event
In this example, we will configure boot.yaml to ignore event logs.

## boot.yaml

```yaml
event:
  - name: my-event
    outputPaths: [""]
echo:
  - name: greeter
    port: 8080
    enabled: true
    eventEntry: my-event
```

## Start main.go
```shell
$ go run main.go

2022-10-19T23:21:22.390+0800    INFO    boot/echo_entry.go:682  Bootstrap EchoEntry     {"eventId": "2aa220d9-343e-4ebb-82f5-f6188098fe42", "entryName": "greeter", "entryType": "EchoEntry"}
```