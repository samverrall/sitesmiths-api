package jwt

import "github.com/gin-gonic/gin"

const (
	CookieName = "__Host-token"
)

type Config struct {
	Insecure bool
}

func (jt *JWTToken) WriteCookie(c *gin.Context, token string) {
	c.SetCookie(CookieName, token, 1000, "/", "", !jt.config.Insecure, true)
}
