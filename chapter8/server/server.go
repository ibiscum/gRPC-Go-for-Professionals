package main

import (
	pb "github.com/ibiscum/gRPC-Go-for-Professionals/chapter8/proto/todov2"
)

type server struct {
	d db

	pb.UnimplementedTodoServiceServer
}
