package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samverrall/sitesmiths-api/internal"
)

type Error struct {
	Message string
	Detail  string
}

func writeError(c *gin.Context, statusCode int, e Error) {
	c.JSON(statusCode, gin.H{
		"error":  e.Message,
		"detail": e.Detail,
	})
}

func writeServiceError(c *gin.Context, err error) bool {
	if err == nil {
		return false
	}

	unwrapped := errors.Unwrap(err)
	body := gin.H{
		"error":  unwrapped.Error(),
		"detail": err.Error(),
	}
	switch {
	case errors.Is(err, internal.ErrBadRequest):
		c.JSON(http.StatusBadRequest, body)
		return true

	case errors.Is(err, internal.ErrNotFound):
		c.JSON(http.StatusNotFound, body)
		return true

	default:
		c.JSON(http.StatusInternalServerError, body)
		return true

	}
}
