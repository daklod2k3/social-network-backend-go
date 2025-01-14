package entity

type LoginEmail struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterEmail struct {
	Name string `json:"name" form:"name" binding:"required"`
	LoginEmail
}

type SessionResponse struct {
	Token string `json:"access_token"`
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
	Name        string `json:"name"`
}
