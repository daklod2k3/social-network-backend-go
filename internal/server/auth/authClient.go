package auth

import (
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
	//Login(form entity.LoginMail) (*pb.AuthResponse, error)
	//Health() (*pb.HealthRes, error)
	//Register(form entity.RegisterMail) (*pb.AuthResponse, error)
}

var (
	address = config.GetConfig().Auth.Url
)

func NewService() *service {

	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("could not connect to auth service:" + err.Error())
	}
	auth := pb.NewAuthClient(conn)

	return &service{conn, auth}
}

func (s *service) CloseGRPC() {
	err := s.client.Close()
	if err != nil {
		log.Fatalln("could not close grpc connection")
	}

}

//
//
//func (s *service) Login(form entity.LoginMail) (*pb.AuthResponse, error) {
//	res, err := s.auth.Login(context.Background(), &pb.LoginEmail{
//		Email:    form.Email,
//		Password: form.Password,
//	})
//	if err != nil {
//		return nil, err
//	}
//	return res, nil
//}
//
//func (s *service) Health() (*pb.HealthRes, error) {
//	res, err := s.auth.Health(context.Background(), nil)
//	if err != nil {
//		return nil, err
//	}
//	return res, nil
//}
//
//
//func (s *service) Register(form entity.RegisterMail) (*pb.AuthResponse, error) {
//	res, err := s.auth.Register(context.Background(), &pb.RegisterEmail{
//		Email:    form.Email,
//		Password: form.Password,
//		Name:     form.Name,
//	})
//	if err != nil {
//		return nil, err
//	}
//	return res, nil
//}
