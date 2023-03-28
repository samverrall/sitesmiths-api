package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/samverrall/sitesmiths-api/internal/site"
)

type API struct {
	port        string
	siteService *site.Service
}

func New(siteSvc *site.Service, port string) *API {
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
		Addr:         ":8080",
		Handler:      r,
		IdleTimeout:  1 * time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return srv
}
