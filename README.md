# Dependencies
## protoc
```
https://github.com/protocolbuffers/protobuf/releases
```

## go pkg
```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go mod tidy
```

# Cmd
```sh 
protoc --go_out=. --go-grpc_out=. proto/*.proto
```

