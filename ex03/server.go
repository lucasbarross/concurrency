package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "grpc-cat/cat"
)

type server struct {
	pb.UnimplementedCatServer
}

func (s *server) Reply(ctx context.Context, in *pb.Package) (*pb.Package, error) {
	return &pb.Package{Message: in.GetMessage()}, nil
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	s := grpc.NewServer()
	pb.RegisterCatServer(s, &server{})
	if err := s.Serve(ln); err != nil {
		log.Fatal(err)
	}
}
