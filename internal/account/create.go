package account

import (
	"context"

	"github.com/google/uuid"
	"github.com/samverrall/sitesmiths-api/account"
	"github.com/samverrall/sitesmiths-api/account/authenticator"
	"github.com/samverrall/sitesmiths-api/internal"
)

type CreatePayload struct {
	Code     string
	Provider string
}

func (s *Service) CreateFromProvider(ctx context.Context, p CreatePayload) error {
	// Check the provider is valid (google etc)
	provider, err := account.NewProvider(p.Provider)
	if err != nil {
		return internal.WrapErr(internal.ErrBadRequest, err)
	}

	// If the provider is valid, try to authenticate to get account details
	// from the provider.
	code, err := authenticator.NewAuthCode(p.Code)
	if err != nil {
		return internal.WrapErr(internal.ErrBadRequest, err)
	}

	token, err := s.authenticator.GetTokenFromCode(ctx, code)
	if err != nil {
		return internal.WrapErr(internal.ErrInternal, err) // TODO: handle this correctly
	}

	accountDetails, err := s.authenticator.GetDetailsFromToken(ctx, token)
	if err != nil {
		return internal.WrapErr(internal.ErrInternal, err) // TODO: handle this correctly
	}

	// Check an account with email doesn't already exist
	_, err = s.repo.GetByEmail(ctx, accountDetails.Email)
	switch {
	case err != nil:
		return internal.WrapErr(internal.ErrInternal, err)
	}

	// Create a new account
	acc := account.New(uuid.New(), accountDetails.Name, accountDetails.Email, provider)

	// Add new account to repo
	if err := s.repo.Add(ctx, acc); err != nil {
		return internal.WrapErr(internal.ErrInternal, err)
	}

	return nil
}
