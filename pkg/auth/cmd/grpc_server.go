package cmd

import (
	"auth/internal"
	"log"
	"net"
	"shared/config"
	"shared/rpc/pb"

	"google.golang.org/grpc"
)

var (
	authPort = config.GetConfig().Auth.Port
)

func StartGRPCServer() {
	lis, err := net.Listen("tcp", ":"+authPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAuthServer(s, internal.NewService())
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
