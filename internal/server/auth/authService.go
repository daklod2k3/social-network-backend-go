package auth

import (
	"auth/entity"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"shared/config"
	"shared/rpc/pb"
)

type service struct {
	client *grpc.ClientConn
	auth   pb.AuthClient
}

type Service interface {
	Login(form entity.LoginEmail) (*pb.AuthResponse, error)
	Register(form entity.RegisterEmail) (*pb.AuthResponse, error)
}

var (
	address = config.GetConfig().Auth.Url
)

func NewService() *service {

	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("could not connect to auth service:" + err.Error())
	}
	auth := pb.NewAuthClient(conn)

	return &service{conn, auth}
}

func (s *service) Login(form entity.LoginEmail) (*pb.AuthResponse, error) {
	res, err := s.auth.Login(context.Background(), &pb.LoginEmail{
		Email:    form.Email,
		Password: form.Password,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *service) Register(form entity.RegisterEmail) (*pb.AuthResponse, error) {
	res, err := s.auth.Register(context.Background(), &pb.RegisterEmail{
		Email:    form.Email,
		Password: form.Password,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *service) CloseGRPC() {
	err := s.client.Close()
	if err != nil {
		fmt.Println("could not close grpc connection")
	}

}
