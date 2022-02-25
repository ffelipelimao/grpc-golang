sudo apt install protobuf-compiler
brew install protobuf 

go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

export PATH="$PATH:$(go env GOPATH)/bin"


protoc --proto_path=proto proto/*.proto --go_out=pb


https://stackoverflow.com/questions/57700860/protoc-gen-go-program-not-found-or-is-not-executable