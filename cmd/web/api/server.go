package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	siteservice "github.com/samverrall/sitesmiths-api/internal/site"
)

type API struct {
	port        string
	siteService *siteservice.Service
}

func New(siteSvc *siteservice.Service, port string) *API {
	return &API{
		siteService: siteSvc,
		port:        port,
	}
}

func (api *API) NewServer() *http.Server {
	r := gin.Default()

	// Site Routes
	sitesRoutes := r.Group("/api/sites")
	sitesRoutes.POST("/", api.CreateSite)
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
