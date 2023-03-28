package site

import (
	"github.com/samverrall/sitesmiths-api/site"
)

type Service struct {
	repo site.Repo
}

func NewService(siteRepo site.Repo) *Service {
	return &Service{
		repo: siteRepo,
	}
}
