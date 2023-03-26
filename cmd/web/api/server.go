package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	siteservice "github.com/samverrall/sitesmiths-api/internal/site"
)

type NewServerArgs struct {
	Port        string
	SiteService *siteservice.Service
}

func NewServer(args NewServerArgs) *http.Server {
	r := gin.Default()

	siteControllers := newSiteControllers(args.SiteService)

	// Site Routes
	sitesRoutes := r.Group("/api/sites")
	sitesRoutes.POST("/", siteControllers.CreateSite)
	///////////////////////

	// User routes
	//userRoutes:= r.Group("/api/users")
	///////////////////////

	// Create a new HTTP server with the Gin router as the handler
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	return srv
}
