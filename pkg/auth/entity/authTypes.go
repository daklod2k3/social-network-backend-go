package entity

import "shared/entity"

type LoginMail struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterMail struct {
	LoginMail
}

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
	*entity.User `json:"user,omitempty"`
}

type SessionRequest struct {
	AccessToken  string `json:"access_token,required"`
	RefreshToken string `json:"refresh_token"`
}

type HealthResponse struct {
	Message string `json:"message"`
}
