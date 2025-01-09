package internal

import (
	"auth/entity"
	"fmt"
	"github.com/supabase-community/gotrue-go/types"
	"github.com/supabase-community/supabase-go"
	"shared/config"
	"shared/utils"
)

type service struct {
	supabase *supabase.Client
	user     *entity.UserRepo
}

type Service interface {
	Login(form entity.LoginEmail) (*types.TokenResponse, error)
	Register(form entity.Register) (*types.SignupResponse, error)
	GetSession()
	Error(err error) entity.SupabaseError
}

func NewService() Service {

	// init supabase supabase
	s := &service{}
	sp := config.GetConfig().Supabase
	client, err := supabase.NewClient(sp.Url, sp.Key, &supabase.ClientOptions{})
	if err != nil {
		panic("supabase supabase init failed" + err.Error())
	}
	s.supabase = client

	s.user = entity.NewRepo()

	return s
}

func (s *service) GetSession() {

}

func (s *service) Login(form entity.LoginEmail) (*types.TokenResponse, error) {
	return s.supabase.Auth.SignInWithEmailPassword(form.Email, form.Password)
}

func (s *service) Register(form entity.Register) (*types.SignupResponse, error) {
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
