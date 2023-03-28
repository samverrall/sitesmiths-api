package account

import (
	"github.com/samverrall/sitesmiths-api/account"
	"github.com/samverrall/sitesmiths-api/account/authenticator"
)

type Service struct {
	repo          account.Repo
	authenticator authenticator.Authenticator
}

func NewService(repo account.Repo, authenticator authenticator.Authenticator) *Service {
	return &Service{
		repo:          repo,
		authenticator: authenticator,
	}
}
