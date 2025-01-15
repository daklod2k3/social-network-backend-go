package auth

import (
	"auth/entity"
	"errors"
	gotrue "github.com/supabase-community/auth-go"
	"github.com/supabase-community/auth-go/types"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"shared/config"

	sharedEntity "shared/entity"
)

type service struct {
	goTrue gotrue.Client
	user   *sharedEntity.UserRepoRead
	*sharedEntity.Service
}

type Service interface {
	Health() (*entity.HealthResponse, error)
	Login(form *entity.LoginMail) (*entity.AuthResponse, error)
	Register(form *entity.RegisterMail) (*entity.AuthResponse, error)
	GetSession(form *entity.SessionRequest) (*entity.AuthResponse, error)
}

func NewService() *service {

	// init supabase supabase
	s := &service{}
	s.Service = sharedEntity.NewService()

	sp := config.GetConfig().Supabase
	client := gotrue.New(sp.Ref, sp.Key)
	_, err := client.GetSettings()
	if err != nil {
		panic("supabase supabase init failed" + err.Error())
	}
	s.goTrue = client

	//init user repo
	s.user = sharedEntity.NewRepoRead(s.Service.Db.GetSchema())

	return s
}

func (s *service) Health() (*entity.HealthResponse, error) {
	_, err := s.goTrue.HealthCheck()
	if err != nil {
		return nil, s.Error(err)
	}

	return &entity.HealthResponse{
		Message: "ok",
	}, nil
}

func (s *service) GetSession(form *entity.SessionRequest) (*entity.AuthResponse, error) {
	cl := s.goTrue.WithToken(form.AccessToken)

	var (
		response = entity.AuthResponse{
			form.AccessToken,
			form.RefreshToken,
			nil,
		}
		user *types.User
	)

	res, err := cl.GetUser()

	if err != nil {
		if form.RefreshToken == "" {
			return nil, s.Error(errors.New("invalid refresh token"))
		}
		tokenRes, err := cl.RefreshToken(form.RefreshToken)
		if err != nil {
			return nil, s.Error(err)
		}
		response.AccessToken = tokenRes.AccessToken
		user = &tokenRes.User
		response.RefreshToken = tokenRes.RefreshToken
	} else {
		user = &res.User
	}

	if user == nil {
		return nil, s.Error(errors.New("invalid token"))
	}

	response.User, _ = s.user.FindUser(&user.ID)
	return &response, nil
}

func (s *service) Login(form *entity.LoginMail) (*entity.AuthResponse, error) {
	res, err := s.goTrue.SignInWithEmailPassword(form.Email, form.Password)
	if err != nil {
		return nil, s.Error(err)
	}
	user, err := s.user.FindUser(&res.User.ID)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			return nil, s.Error(err)
		}
	}
	return &entity.AuthResponse{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
		User:         user,
	}, nil
}

func (s *service) Register(form *entity.RegisterMail) (*entity.AuthResponse, error) {
	auth, err := s.goTrue.Signup(types.SignupRequest{
		Email:    form.Email,
		Password: form.Password,
	})
	if err != nil {
		return nil, s.Error(err)
	}

	return &entity.AuthResponse{
		AccessToken:  auth.AccessToken,
		RefreshToken: auth.RefreshToken,
	}, nil

}

func (s *service) Error(err error) error {
	s.Service.Logger.Error(err.Error())
	return status.Error(codes.Aborted, err.Error())
}
