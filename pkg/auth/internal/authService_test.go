package internal

import (
	"context"
	"shared/rpc/pb"
	"testing"
)

var (
	loginEmail    = "test@test.com"
	loginPass     = "123123"
	registerEmail = "test@test.com"
	registerPass  = "123123"
)

func Test(t *testing.T) {
	s := NewService()
	ctx := context.Background()

	t.Log(s.Login(ctx, &pb.LoginEmail{
		Email:    loginEmail,
		Password: loginPass,
	}))

}
