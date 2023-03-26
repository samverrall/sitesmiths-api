package internal

import "github.com/samverrall/sitesmiths-api/repo"

type SiteService struct {
	repo repo.Site
}

func NewSiteService(siteRepo repo.Site) *SiteService {
	return &SiteService{
		repo: siteRepo,
	}
}
