package authenticator

import (
	"context"
	"errors"

	"github.com/samverrall/sitesmiths-api/account"
)

var (
	ErrAuthCodeFailure       = errors.New("failed to use auth code to get token")
	ErrAccountDetailsFailure = errors.New("failed to get account details from token")
)

type Authenticator interface {
	GetTokenFromCode(ctx context.Context, code AuthCode) (Token, error)
	GetDetailsFromToken(ctx context.Context, token Token) (*account.Account, error)
}
