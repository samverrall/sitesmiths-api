package account

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/samverrall/sitesmiths-api/account"
	"github.com/samverrall/sitesmiths-api/account/authenticator"
	"github.com/samverrall/sitesmiths-api/internal"
)

var (
	ErrAccountExists = errors.New("account already exists")
)

// CreateFromProviderPayload defines a primative seriaisable payload with fields
// required for the CreateFromProvider method.
type CreateFromProviderPayload struct {
	Code     string
	Provider string
}

func (s *Service) CreateFromProvider(ctx context.Context, p CreateFromProviderPayload) error {
	// Check the provider is valid (google etc)
	provider, err := account.NewProvider(p.Provider)
	if err != nil {
		return errors.Join(internal.ErrBadRequest, err)
	}

	code, err := authenticator.NewAuthCode(p.Code)
	if err != nil {
		return errors.Join(internal.ErrBadRequest, err)
	}

	// If the provider is valid and we have a valid auth code, try to authenticate
	// to get account details from the provider.
	token, err := s.authenticator.GetTokenFromCode(ctx, code)
	if err != nil {
		return errors.Join(internal.ErrInternal, err)
	}

	accountDetails, err := s.authenticator.GetDetailsFromToken(ctx, token)
	if err != nil {
		return errors.Join(internal.ErrInternal, err)
	}

	// Check an account with email doesn't already exist
	existingAccount, err := s.repo.GetByEmail(ctx, accountDetails.Email)
	switch {
	case err != nil && !errors.Is(err, account.ErrNotFound):
		return errors.Join(internal.ErrInternal, err)

	case existingAccount.Active:
		return errors.Join(internal.ErrInternal, ErrAccountExists)
	}

	// Create a new account
	acc := account.New(uuid.New(), accountDetails.Name, accountDetails.Email, provider)

	// Add new account to repo
	if err := s.repo.Add(ctx, acc); err != nil {
		return errors.Join(internal.ErrInternal, err)
	}

	return nil
}
