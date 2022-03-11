```sudo apt install protobuf-compiler```

```brew install protobuf```

```go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26```
```go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1```
```go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest```

```export PATH="$PATH:$(go env GOPATH)/bin"```

```protoc --proto_path=proto proto/*.proto --go_out=pb```
```protoc --proto_path=proto proto/*.proto --go-grpc_out=pb```

```$ go install github.com/ktr0731/evans@latest```


https://stackoverflow.com/questions/57700860/protoc-gen-go-program-not-found-or-is-not-executable
https://stackoverflow.com/questions/60578892/protoc-gen-go-grpc-program-not-found-or-is-not-executable