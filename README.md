# protoc-gen-tars


## Install

```
go get github.com/go-tars/protoc-gen-tars@master
```

Also required:

- [protoc](https://github.com/google/protobuf)
- [protoc-gen-go](https://github.com/golang/protobuf)

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
protoc --gogofaster_out=. --tars_out=. example.proto
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