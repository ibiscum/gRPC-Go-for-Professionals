// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: proto/dummy/dummy.proto

package dummy

import (
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const ()

// DummyServiceClient is the client API for DummyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DummyServiceClient interface {
}

type dummyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDummyServiceClient(cc grpc.ClientConnInterface) DummyServiceClient {
	return &dummyServiceClient{cc}
}

// DummyServiceServer is the server API for DummyService service.
// All implementations must embed UnimplementedDummyServiceServer
// for forward compatibility
type DummyServiceServer interface {
	mustEmbedUnimplementedDummyServiceServer()
}

// UnimplementedDummyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDummyServiceServer struct {
}

func (UnimplementedDummyServiceServer) mustEmbedUnimplementedDummyServiceServer() {}

// UnsafeDummyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DummyServiceServer will
// result in compilation errors.
type UnsafeDummyServiceServer interface {
	mustEmbedUnimplementedDummyServiceServer()
}

func RegisterDummyServiceServer(s grpc.ServiceRegistrar, srv DummyServiceServer) {
	s.RegisterService(&DummyService_ServiceDesc, srv)
}

// DummyService_ServiceDesc is the grpc.ServiceDesc for DummyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DummyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dummy.DummyService",
	HandlerType: (*DummyServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "proto/dummy/dummy.proto",
}
