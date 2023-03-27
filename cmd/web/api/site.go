package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	siteservice "github.com/samverrall/sitesmiths-api/internal/site"
)

func (a *API) CreateSite(c *gin.Context) {
	ctx := c.Request.Context()

	var payload struct {
		Name    string `json:"name"`
		URL     string `json:"url"`
		OwnerID string `json:"ownerId"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		writeError(c, http.StatusBadRequest, Error{
			Message: "Invalid site payload supplied",
			Detail:  err.Error(),
		})
		return
	}

	err := a.siteService.Create(ctx, siteservice.CreatePayload{
		Name:    payload.Name,
		URL:     payload.URL,
		OwnerID: payload.OwnerID,
	})
	if stop := writeServiceError(c, err); stop {
		return
	}

	c.Status(http.StatusCreated)
}
