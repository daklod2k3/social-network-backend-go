package cmd

import (
	"log"
	"net"
	"os"
	"shared/rpc/pb"

	"google.golang.org/grpc"
)

var (
	auth_port = os.Getenv("AUTH_PORT")
)

type server struct {
	pb.UnimplementedAuthServer
}

func StartGRPCServer() {
	lis, err := net.Listen("tcp", auth_port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAuthServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
