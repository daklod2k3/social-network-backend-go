package auth

import (
	"errors"
	gotrue "github.com/supabase-community/auth-go"
	"github.com/supabase-community/auth-go/types"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	authEntity "shared/entity/auth"
	"shared/global"

	sharedEntity "shared/entity"
)

var (
	logger = global.Logger
)

type service struct {
	goTrue gotrue.Client
	user   *sharedEntity.UserRepoRead
	*sharedEntity.Service
}

func NewService() *service {

	// init supabase supabase
	s := &service{}
	s.Service = sharedEntity.NewService()

	sp := global.Config.Supabase
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

func (s *service) Health() (*authEntity.HealthResponse, error) {
	_, err := s.goTrue.HealthCheck()
	if err != nil {
		return nil, s.Error(err)
	}

	return &authEntity.HealthResponse{
		Message: "ok",
	}, nil
}

func (s *service) GetSession(form *authEntity.SessionRequest) (*authEntity.AuthResponse, error) {

	cl := s.goTrue.WithToken(form.AccessToken)

	var (
		response = authEntity.AuthResponse{
			form.AccessToken,
			form.RefreshToken,
			nil,
			nil,
		}
		user *types.User
	)

	logger.Info("Session: " + form.AccessToken + form.RefreshToken)

	res, err := cl.GetUser()

	if err != nil {
		logger.Error(err.Error())
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
		response.UserId = &tokenRes.User.ID
	} else {
		user = &res.User
		response.UserId = &res.User.ID
	}

	if user == nil {
		return nil, s.Error(errors.New("invalid token"))
	}

	response.User, _ = s.user.FindUser(&user.ID)
	return &response, nil
}

func (s *service) Login(form *authEntity.LoginMail) (*authEntity.AuthResponse, error) {
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
	return &authEntity.AuthResponse{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
		User:         user,
		UserId:       &res.User.ID,
	}, nil
}

func (s *service) Register(form *authEntity.RegisterMail) (*authEntity.AuthResponse, error) {
	auth, err := s.goTrue.Signup(types.SignupRequest{
		Email:    form.Email,
		Password: form.Password,
	})
	if err != nil {
		return nil, s.Error(err)
	}

	return &authEntity.AuthResponse{
		AccessToken:  auth.AccessToken,
		RefreshToken: auth.RefreshToken,
		UserId:       &auth.User.ID,
	}, nil

}

func (s *service) Error(err error) error {
	s.Service.Logger.Error(err.Error())
	return status.Error(codes.Aborted, err.Error())
}
