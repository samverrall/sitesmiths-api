package google

import (
	"context"

	"github.com/samverrall/sitesmiths-api/account"
	"github.com/samverrall/sitesmiths-api/account/authenticator"
)

var (
	_ authenticator.Authenticator = &Authenticator{}
)

type Authenticator struct {
}

func NewAuthenticator() *Authenticator {
	return &Authenticator{}
}

func (a *Authenticator) GetTokenFromCode(ctx context.Context, code authenticator.AuthCode) (authenticator.Token, error) {
	return "", nil
}

func (a *Authenticator) GetDetailsFromToken(ctx context.Context, token authenticator.Token) (*account.Account, error) {
	return nil, nil
}
