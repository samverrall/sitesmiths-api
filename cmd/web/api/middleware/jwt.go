package middleware

import (
	"context"
	"errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	apijwt "github.com/samverrall/sitesmiths-api/cmd/web/api/jwt"

	"github.com/golang-jwt/jwt"
)

const (
	UserContextTokenKey = "user"
)

func JWTAuthenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		token, err := c.Request.Cookie(apijwt.CookieName)
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				respondWithError(c, err, http.StatusUnauthorized, "Access token required")
			}

			respondWithError(c, err, http.StatusInternalServerError, "Failed to authenticate user")
		}

		t, err := jwt.ParseWithClaims(token.String(), &apijwt.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SPACECMS_SECRET")), nil
		})
		if stop := respondWithError(c, err, http.StatusInternalServerError, "Failed to parse JWT"); stop {
			return
		}

		claims, ok := t.Claims.(*apijwt.Claims)
		if !ok || !t.Valid {
			err := errors.New("failed to read token claims")
			respondWithError(c, err, http.StatusInternalServerError, err.Error())
		}

		// Append the user to the context
		ctx = context.WithValue(ctx, UserContextTokenKey, claims)

		c.Next()
	}
}

func ClaimsFromContext(ctx context.Context) (*apijwt.Claims, error) {
	claims, ok := ctx.Value(UserContextTokenKey).(*apijwt.Claims)
	if !ok {
		return nil, errors.New("failed to get claims from ctx")
	}
	return claims, nil
}
