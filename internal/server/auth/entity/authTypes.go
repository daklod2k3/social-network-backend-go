package entity

type LoginEmail struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Register struct {
	Name string `json:"name" form:"name" binding:"required"`
	LoginEmail
}

type SessionResponse struct {
	Token string `json:"access_token"`
}
