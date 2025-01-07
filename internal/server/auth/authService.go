package auth

import (
	"core/internal/server/auth/entity"
	"core/shared/utils"
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"github.com/supabase-community/gotrue-go/types"
	"github.com/supabase-community/supabase-go"
)

type service struct {
	supabase *supabase.Client
	user     entity.UserRepo
}

type Service interface {
	Login()
	Register()
	GetSession()
}

func NewService() *service {

	// init supabase supabase
	s := &service{}
	url := viper.GetString("supabase_url")
	key := viper.GetString("supabase_key")
	client, err := supabase.NewClient(url, key, &supabase.ClientOptions{})
	if err != nil {
		panic("supabase supabase init failed" + err.Error())
	}
	s.supabase = client
	
	return s
}

func (s *service) Login(form entity.LoginEmail) (*types.TokenResponse, error) {
	return s.supabase.Auth.SignInWithEmailPassword(form.Email, form.Password)
}

func (s *service) Register(form entity.Register) (*types.SignupResponse, error) {
	//auth, err := s.supabase.Auth.Signup(types.SignupRequest{
	//	Email:    form.Email,
	//	Password: form.Password,
	//	Data: map[string]interface{}{
	//		"display_name": form.Name,
	//	},
	//})
	//if err != nil {
	//	return nil, err
	//}

	auth := &types.SignupResponse{
		User: types.User{
			ID: uuid.New(),
		},
	}

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

type SupabaseError struct {
	Code      int    `json:"code"`
	ErrorCode string `json:"error_code"`
	Msg       string `json:"msg"`
}

func (s *service) Error(err error) SupabaseError {
	var parse SupabaseError
	e := utils.Deserialize(err.Error(), &parse)
	if e != nil {
		// supabase error convert
		fmt.Println(e)
	}
	return parse
}
