package interfaces

import (
	authEntity "shared/entity/auth"
)

type AuthService interface {
	GetSession(request *authEntity.SessionRequest) (*authEntity.AuthResponse, error)
	Login(mail *authEntity.LoginMail) (*authEntity.AuthResponse, error)
	Register(mail *authEntity.RegisterMail) (*authEntity.AuthResponse, error)
	Health() (*authEntity.HealthResponse, error)
}

type AuthRpc interface {
}
