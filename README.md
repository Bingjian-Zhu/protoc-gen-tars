# protoc-gen-tars


## Install

```
go get github.com/bingjian-zhu/protoc-gen-tars@master
```

Also required:

- [protoc](https://github.com/google/protobuf)
- [protoc-gen-go](https://github.com/golang/protobuf)

注意：protoc-gen-go版本要选v1.3.0
```
go build .
```

## Usage

Define your service as example.proto

```
syntax = "proto3";

service Greeter {
	rpc Hello(Request) returns (Response) {}
}

message Request {
	string name = 1;
}

message Response {
	string msg = 1;
}
```

Generate the code:

```
protoc --go_out=. --tars_out=. example.proto
```

Maybe gogofaster is better:

```
protoc --go_out=. --pb2tars_out=. example.proto
```

Your output result should be:

```
./
├── example.pb.go
├── example.pb.tars.go
└── example.proto
```

## Example test

server:
```
go run example/server/main.go -config example/config.conf
```

client:
```
go run example/client/main.go -config example/config.conf
```

output:
```
result is: hellosandyskies
```