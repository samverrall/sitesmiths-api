package page

import "github.com/samverrall/sitesmiths-api/page"

type Service struct {
	pageRepo page.Repo
}

func NewService(pageRepo page.Repo) *Service {
	return &Service{
		pageRepo: pageRepo,
	}
}
