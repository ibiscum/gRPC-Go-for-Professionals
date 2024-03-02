module github.com/ibiscum/gRPC-Go-for-Professionals/client

go 1.20

replace github.com/ibiscum/gRPC-Go-for-Professionals/proto => ../proto

require (
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.1
	google.golang.org/grpc v1.59.0
	google.golang.org/protobuf v1.31.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	golang.org/x/net v0.18.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231106174013-bbf56f31fb17 // indirect
)
