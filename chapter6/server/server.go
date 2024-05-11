package main

import (
	pb "github.com/ibiscum/gRPC-Go-for-Professionals/chapter6/proto/todo/v2"
)

type server struct {
	d db

	pb.UnimplementedTodoServiceServer
}
