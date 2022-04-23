# go-grpc

### Go gRPC for macOS
1) Install Protobuf on macOS
```
brew install protobuf
```
2) Install Evans gRPC client for macOS
```
brew tap ktr0731/evans
brew install evans
```
3) Install vscode-proto3 for VSCode https://marketplace.visualstudio.com/items?itemName=zxh404.vscode-proto3

4) Go get package in project
```
go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
```
5) Install gRPC tool in project
```
go install google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```
