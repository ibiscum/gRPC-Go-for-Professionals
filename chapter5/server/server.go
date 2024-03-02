package main

import (
	pb "github.com/ibiscum/gRPC-Go-for-Professionals/chapter5/proto/todo/v1"
)

type server struct {
	d db

	pb.UnimplementedTodoServiceServer
}
