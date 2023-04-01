package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samverrall/sitesmiths-api/internal/account"
)

func (a *API) CreateAccountFromProvider(c *gin.Context) {
	ctx := c.Request.Context()

	var payload struct {
		Code     string `json:"code"`
		Provider string `json:"provider"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		writeError(c, http.StatusBadRequest, Error{
			Message: "Invalid site payload supplied",
			Detail:  err.Error(),
		})
		return
	}

	err := a.accountService.CreateFromProvider(ctx, account.CreateFromProviderPayload{
		Provider: payload.Code,
		Code:     payload.Code,
	})
	if stop := writeServiceError(c, err); stop {
		return
	}

	c.Status(http.StatusCreated)
	return
}
