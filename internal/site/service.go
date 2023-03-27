package siteservice

import (
	"github.com/samverrall/sitesmiths-api/site"
)

type Service struct {
	repo site.Repo
}

func New(siteRepo site.Repo) *Service {
	return &Service{
		repo: siteRepo,
	}
}
