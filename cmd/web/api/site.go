package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samverrall/sitesmiths-api/cmd/web/api/middleware"
	"github.com/samverrall/sitesmiths-api/internal/site"
)

func (a *API) CreateSite(c *gin.Context) {
	ctx := c.Request.Context()

	claims, err := middleware.ClaimsFromContext(ctx)
	if writeError(c, http.StatusInternalServerError, err, "Failed to get user from token claims") {
		return
	}

	var payload struct {
		Name    string `json:"name"`
		URL     string `json:"url"`
		OwnerID string `json:"ownerId"`
	}
	payload.OwnerID = claims.UserUUID
	err = c.ShouldBindJSON(&payload)
	if writeError(c, http.StatusBadRequest, err, "Invalid site payload supplied") {
		return
	}

	err = a.siteService.Create(ctx, site.CreatePayload{
		Name:    payload.Name,
		URL:     payload.URL,
		OwnerID: payload.OwnerID,
	})
	if writeServiceError(c, err) {
		return
	}

	c.Status(http.StatusCreated)
}
