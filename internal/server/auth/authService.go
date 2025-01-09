package auth

import (
	"auth/entity"
	"context"
	"google.golang.org/grpc"
	"os"
	"shared/rpc/pb"
)

type service struct {
	client pb.AuthClient
}

type Service interface {
	Login(form entity.LoginEmail) (*pb.LoginResponse, error)
}

var (
	address = os.Getenv("auth_url")
)

func NewService() *service {

	conn, err := grpc.NewClient(address)
	if err != nil {
		panic("could not connect to auth service:" + err.Error())
	}
	defer conn.Close()

	c := pb.NewAuthClient(conn)

	return &service{c}
}

func (s *service) Login(form entity.LoginEmail) (*pb.LoginResponse, error) {
	res, err := s.client.Login(context.Background(), &pb.LoginRequest{
		Email:    form.Email,
		Password: form.Password,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
