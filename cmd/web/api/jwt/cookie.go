package jwt

import "github.com/gin-gonic/gin"

const (
	CookieName = "__Host-token"
	MaxAge     = 1000
	HttpOnly   = true
)

type Config struct {
	Insecure bool
}

func (jt *JWTToken) WriteCookie(c *gin.Context, token string) {
	c.SetCookie(CookieName, token, MaxAge, "/", "", !jt.config.Insecure, HttpOnly)
}
