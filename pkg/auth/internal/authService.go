package internal

import (
	"auth/entity"
	"fmt"
	"github.com/supabase-community/gotrue-go/types"
	"github.com/supabase-community/supabase-go"
	"golang.org/x/net/context"
	"shared/config"
	"shared/rpc/pb"
)

type Service struct {
	supabase *supabase.Client
	user     *entity.UserRepo
	pb.UnimplementedAuthServer
}

type service interface {
	Login(form pb.LoginEmail) pb.AuthResponse
	//Register(form pb.LoginEmail) pb.AuthResponse
	//GetSession()
}

func NewService() *Service {

	// init supabase supabase
	s := &Service{}
	sp := config.GetConfig().Supabase
	client, err := supabase.NewClient(sp.Url, sp.Key, &supabase.ClientOptions{})
	if err != nil {
		panic("supabase supabase init failed" + err.Error())
	}
	s.supabase = client

	s.user = entity.NewRepo()

	return s
}

func (s *Service) GetSession() {

}

func (s *Service) Login(_ context.Context, form *pb.LoginEmail) (*pb.AuthResponse, error) {
	res, err := s.supabase.Auth.SignInWithEmailPassword(form.Email, form.Password)
	if err != nil {
		return &pb.AuthResponse{}, err
	}
	return &pb.AuthResponse{
		AccessToken: res.AccessToken,
	}, err
}

func (s *Service) Register(form entity.Register) (*types.SignupResponse, error) {
	auth, err := s.supabase.Auth.Signup(types.SignupRequest{
		Email:    form.Email,
		Password: form.Password,
		Data: map[string]interface{}{
			"display_name": form.Name,
		},
	})
	if err != nil {
		return nil, err
	}

	//auth := &types.SignupResponse{
	//	User: types.User{
	//		ID: uuid.New(),
	//	},
	//}

	user, err := s.user.CreateUser(&auth.ID, &form.Name)

	fmt.Println(user)

	if err != nil {
		s.supabase.Auth.AdminDeleteUser(types.AdminDeleteUserRequest{
			UserID: user.UserId,
		})
		return nil, err
	}
	return auth, nil

}
