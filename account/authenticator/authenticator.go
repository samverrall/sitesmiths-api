package authenticator

import (
	"context"

	"github.com/samverrall/sitesmiths-api/account"
)

type Token string

type Authenticator interface {
	GetTokenFromCode(ctx context.Context, code AuthCode) (Token, error)
	GetDetailsFromToken(ctx context.Context, token Token) (*account.Account, error)
}
