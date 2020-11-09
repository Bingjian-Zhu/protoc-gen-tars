# protoc-gen-go


## build
```bash
go build .

cp protoc-gen-go $GOPATH/bin
```

## Use

```bash
protoc --go_out=plugins=tarsrpc:. example.proto
```