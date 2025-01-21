package interfaces

import (
	"github.com/supabase-community/auth-go"
)

type Server interface {
	GetAuthClient() *auth.Client
}
