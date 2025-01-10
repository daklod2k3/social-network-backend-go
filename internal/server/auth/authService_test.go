package auth

import "testing"

var (
	loginEmail    = "test@test.com"
	loginPass     = "123123"
	registerEmail = "test@test.com"
	registerPass  = "123123"
)

func Test(t *testing.T) {
	s := NewService()
	s.Login()
}
