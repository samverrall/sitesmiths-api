package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	siteservice "github.com/samverrall/sitesmiths-api/internal/site"
)

type SiteControllers struct {
	siteService *siteservice.Service
}

func newSiteControllers(siteService *siteservice.Service) *SiteControllers {
	return &SiteControllers{
		siteService: siteService,
	}
}

func (s *SiteControllers) CreateSite(c *gin.Context) {
	c.Status(http.StatusAccepted)
}
