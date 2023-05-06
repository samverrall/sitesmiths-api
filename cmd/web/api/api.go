package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/samverrall/sitesmiths-api/cmd/web/api/jwt"
	"github.com/samverrall/sitesmiths-api/cmd/web/api/middleware"
	"github.com/samverrall/sitesmiths-api/internal/account"
	"github.com/samverrall/sitesmiths-api/internal/site"
)

type API struct {
	insecure       bool
	port           string
	siteService    *site.Service
	accountService *account.Service
	jwtTokens      *jwt.JWTToken
}

func New(siteSvc *site.Service, accountSvc *account.Service, port string, insecure bool) *API {
	return &API{
		siteService:    siteSvc,
		accountService: accountSvc,
		port:           port,
		jwtTokens: jwt.New(10, "access-token-secret", "refresh-token-secret", &jwt.Config{
			Insecure: insecure,
		}),
	}
}

func (api *API) NewServer() *http.Server {
	r := gin.Default()

	r.Use(middleware.JWTAuthenticate())

	// Site Routes
	sitesRoutes := r.Group("/api/sites")
	sitesRoutes.POST("/", api.CreateSite)
	///////////////////////

	// User routes
	userRoutes := r.Group("/api/users")
	userRoutes.POST("/", api.CreateAccountFromProvider)
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
