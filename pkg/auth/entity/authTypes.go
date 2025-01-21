package entity

type LoginMail struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterMail struct {
	LoginMail
}
