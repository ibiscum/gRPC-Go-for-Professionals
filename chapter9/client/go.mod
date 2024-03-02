module github.com/ibiscum/gRPC-Go-for-Professionals/chapter9/client

go 1.20

replace github.com/ibiscum/gRPC-Go-for-Professionals/chapter9/proto => ../proto

require (
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.1
	github.com/ibiscum/gRPC-Go-for-Professionals/proto v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.62.0
	google.golang.org/protobuf v1.32.0
)

require (
	github.com/envoyproxy/protoc-gen-validate v1.0.4 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	golang.org/x/net v0.21.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240123012728-ef4313101c80 // indirect
)
