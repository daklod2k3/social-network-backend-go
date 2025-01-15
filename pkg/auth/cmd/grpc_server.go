package cmd

import (
	"auth/entity"
	"auth/internal/auth"
	"context"
	"fmt"
	"log"
	"net"
	"shared/config"
	"shared/rpc/pb"

	"google.golang.org/grpc"
)

var (
	grpcPort = config.GetConfig().Auth.Grpc.Port
)

type server struct {
	pb.UnimplementedAuthServer
	authService auth.Service
}

func StartGRPCServer() {
	lis, err := net.Listen("tcp", ":"+fmt.Sprint(grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAuthServer(s, server{})
	log.Printf("GRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s server) GetSession(_ context.Context, req *pb.SessionReq) (*pb.AuthResponse, error) {
	session, err := s.authService.GetSession(&entity.SessionRequest{
		AccessToken: req.AccessToken,
	})
	if err != nil {
		return nil, err
	}
	return &pb.AuthResponse{
		AccessToken: session.AccessToken,
	}, nil
}

//func (s server) GetProfile(_ context.Context, req *pb.)
