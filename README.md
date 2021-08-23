# Greet Service
This is a minimal project using all 4 patterns of gRPC call.

## Requirements
- `Go >= 1.16.7`
- `protoc >= 3.17.3`
- `protoc-gen-go >= 1.26`
- `protoc-gen-go-grpc >= 1.1`

Make sure you've append path to `go`, `protoc` and `go env GOPATH` in your `PATH` enviroment variable!

## Quickstart
Generating protobuf and gRPC files for Go:
```=
make grpc
```
If you did any change to `.proto` files, make sure you call this first.

Start server:
```=
go run server/*
```

In other shell, start client:
```=
go run client/*
# or `make client`
```

## License
This repository is distributed under MIT license.
