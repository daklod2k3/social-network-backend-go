package interfaces

import (
	"auth/entity"
	authEntity "shared/entity/auth"
)

type AuthService interface {
	GetSession(request *authEntity.SessionRequest) (*authEntity.AuthResponse, error)
	Login(mail *entity.LoginMail) (*authEntity.AuthResponse, error)
	Register(mail *entity.RegisterMail) (*authEntity.AuthResponse, error)
	Health() (*authEntity.HealthResponse, error)
}

type AuthRpc interface {
}
