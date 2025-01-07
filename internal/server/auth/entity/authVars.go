package entity

type LoginEmail struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Register struct {
	Name string `json:"name" form:"name" binding:"required"`
	LoginEmail
}
