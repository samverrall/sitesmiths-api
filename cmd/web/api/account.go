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
	err := c.ShouldBindJSON(&payload)
	if stop := writeError(c, http.StatusBadRequest, err, "Invalid site payload supplied"); stop {
		return
	}

	accountID, err := a.accountService.CreateFromProvider(ctx, account.CreateFromProviderPayload{
		Provider: payload.Code,
		Code:     payload.Code,
	})
	if stop := writeServiceError(c, err); stop {
		return
	}

	token, err := a.jwtTokens.CreateTokenPair(ctx, accountID)
	if stop := writeError(c, http.StatusInternalServerError, err, "Failed to create token pair"); stop {
		return
	}

	a.jwtTokens.WriteCookie(c, token.AccessToken)
	c.Status(http.StatusCreated)
	return
}
