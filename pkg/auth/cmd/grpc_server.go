package cmd

import (
	"auth/internal/auth"
	"context"
	"fmt"
	"log"
	"net"
	"shared/config"
	authEntity "shared/entity/auth"
	"shared/interfaces"
	"shared/rpc/pb"

	"google.golang.org/grpc"
)

var (
	grpcPort = config.GetConfig().Auth.Grpc.Port
)

type server struct {
	pb.UnimplementedAuthServer
	authService interfaces.AuthService
}

func StartGRPCServer() {
	lis, err := net.Listen("tcp", ":"+fmt.Sprint(grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAuthServer(s, server{
		authService: auth.NewService(),
	})
	log.Printf("GRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s server) GetSession(_ context.Context, req *pb.SessionReq) (*pb.AuthResponse, error) {
	session, err := s.authService.GetSession(&authEntity.SessionRequest{
		AccessToken: req.AccessToken,
	})
	if err != nil {
		return nil, err
	}
	var userPp pb.User
	if session.User != nil {
		userPp = pb.User{
			Id:          session.User.ID.String(),
			DisplayName: session.User.DisplayName,
			AvatarPath:  session.User.AvatarPath,
			UserId:      session.User.ID.String(),
		}
	}
	return &pb.AuthResponse{
		AccessToken:  session.AccessToken,
		RefreshToken: session.RefreshToken,
		UserId:       session.UserId.String(),
		User:         &userPp,
	}, nil
}

//func (s server) GetProfile(_ context.Context, req *pb.)
