package account

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/samverrall/sitesmiths-api/account"
	"github.com/samverrall/sitesmiths-api/account/authenticator"
)

func TestCreateFromProvider(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	ctx := context.Background()

	authMock := authenticator.NewMockAuthenticator(ctrl)
	repo := account.NewMockRepo(ctrl)

	token := authenticator.Token("token")
	validCode := authenticator.AuthCode("valid")
	validEmail := account.Email("test@test.com")

	svc := NewService(repo, authMock)

	t.Run("invalid provider returns invalid provider error", func(t *testing.T) {
		err := svc.CreateFromProvider(ctx, CreatePayload{
			Provider: "invalid",
			Code:     validCode.String(),
		})
		if err != nil && !errors.Is(err, account.ErrInvalidProvider) {
			t.Errorf("want ErrInvalidProvider error got: %v", err)
			return
		}
		if err == nil {
			t.Error("want ErrInvalidProvider got <nil>")
			return
		}
	})

	t.Run("empty code returns empty code error", func(t *testing.T) {
		err := svc.CreateFromProvider(ctx, CreatePayload{
			Provider: account.GoogleProvider.String(),
			Code:     "",
		})
		if err != nil && !errors.Is(err, authenticator.ErrEmptyAuthCode) {
			t.Errorf("want ErrEmptyAuthCode error got: %v", err)
			return
		}
		if err == nil {
			t.Error("want ErrEmptyAuthCode got <nil>")
			return
		}

	})

	t.Run("account successfully creates from code", func(t *testing.T) {
		repo.EXPECT().Add(ctx, gomock.Any()).Return(nil)
		repo.EXPECT().GetByEmail(ctx, validEmail).Return(account.Account{}, account.ErrNotFound)
		authMock.EXPECT().GetTokenFromCode(ctx, validCode).Return(token, nil)
		acc := account.Account{
			Email: validEmail,
			Name:  "test",
		}
		authMock.EXPECT().GetDetailsFromToken(ctx, token).Return(&acc, nil)

		err := svc.CreateFromProvider(ctx, CreatePayload{
			Provider: account.GoogleProvider.String(),
			Code:     validCode.String(),
		})
		if err != nil {
			t.Errorf("want <nil> error got: %v", err)
		}
	})

	t.Run("account that already exists returns an exists error", func(t *testing.T) {
		repo.EXPECT().GetByEmail(ctx, validEmail).Return(account.Account{
			Active: true,
		}, nil)
		authMock.EXPECT().GetTokenFromCode(ctx, validCode).Return(token, nil)
		acc := account.Account{
			Email: validEmail,
			Name:  "test",
		}
		authMock.EXPECT().GetDetailsFromToken(ctx, token).Return(&acc, nil)

		err := svc.CreateFromProvider(ctx, CreatePayload{
			Provider: account.GoogleProvider.String(),
			Code:     validCode.String(),
		})
		if err != nil && !errors.Is(err, ErrAccountExists) {
			t.Errorf("want <nil> error got: %v", err)
		}
		if err == nil {
			t.Error("want ErrAccountExists got <nil>")
		}
	})
}
