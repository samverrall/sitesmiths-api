package google

import (
	"context"
	"errors"
	"fmt"

	"github.com/samverrall/sitesmiths-api/account"
	"github.com/samverrall/sitesmiths-api/account/authenticator"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	oAuthV2 "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

var (
	_ authenticator.Authenticator = &Authenticator{}
)

type Authenticator struct {
	clientID     string
	clientSecret string
	redirectURL  string
}

func NewAuthenticator(clientID, clientSecret, redirectURL string) *Authenticator {
	return &Authenticator{
		clientID:     clientID,
		clientSecret: clientSecret,
		redirectURL:  redirectURL,
	}
}

func (a *Authenticator) GetTokenFromCode(ctx context.Context, code authenticator.AuthCode) (authenticator.Token, error) {
	config := &oauth2.Config{
		ClientID:     a.clientID,
		ClientSecret: a.clientSecret,
		RedirectURL:  a.redirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	token, err := config.Exchange(ctx, code.String())
	if err != nil {
		return "", errors.Join(authenticator.ErrAuthCodeFailure, err)
	}

	return authenticator.Token(token.AccessToken), nil
}

func (a *Authenticator) GetDetailsFromToken(ctx context.Context, token authenticator.Token) (*account.Account, error) {

	client := oauth2.NewClient(ctx, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token.String()},
	))

	srv, err := oAuthV2.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, fmt.Errorf("failed to create oauth2 service: %v", err)
	}

	userinfoService := oAuthV2.NewUserinfoService(srv)
	userinfo, err := userinfoService.Get().Do()
	if err != nil {
		return nil, errors.Join(authenticator.ErrAccountDetailsFailure, err)
	}

	return &account.Account{
		Name:  account.Name(userinfo.Name),
		Email: account.Email(userinfo.Email),
	}, nil
}
