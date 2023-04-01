package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samverrall/sitesmiths-api/internal/site"
)

func (a *API) CreateSite(c *gin.Context) {
	ctx := c.Request.Context()

	var payload struct {
		Name    string `json:"name"`
		URL     string `json:"url"`
		OwnerID string `json:"ownerId"`
	}
	err := c.ShouldBindJSON(&payload)
	if writeError(c, http.StatusBadRequest, Error{
		Err:     err,
		Message: "Invalid site payload supplied",
		Detail:  err.Error(),
	}) {
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
