package middleware

import "github.com/gin-gonic/gin"

func respondWithError(c *gin.Context, err error, code int, message string) bool {
	if err == nil {
		return false
	}

	c.AbortWithStatusJSON(code, gin.H{"error": message})
	return true
}
