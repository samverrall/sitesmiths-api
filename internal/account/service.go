package account

import "github.com/samverrall/sitesmiths-api/account"

type Service struct {
	repo account.Repo
}

func NewService(repo account.Repo) *Service {
	return &Service{
		repo: repo,
	}
}
