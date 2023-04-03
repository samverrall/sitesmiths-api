package jwt

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type Token struct {
	UserUUID               string `json:"user_uuid"`
	AccessToken            string `json:"access_token"`
	RefreshToken           string `json:"refresh_token"`
	AccessExpiryTime       int64  `json:"access_expiry_time"`
	RefreshExpiryTime      int64  `json:"refresh_expiry_time"`
	AccessExpiryTimeStamp  string `json:"access_expiry_time_stamp"`
	RefreshExpiryTimeStamp string `json:"refresh_expiry_time_stamp"`
}

type JWTToken struct {
	accessTokenExpiryMinutes int
	secretAT                 string
	secretRT                 string
	config                   *Config
}

func New(accessTokenExpiryMinutes int, accessTokenSecret, refreshTokenSecret string, c *Config) *JWTToken {
	return &JWTToken{
		accessTokenExpiryMinutes: accessTokenExpiryMinutes,
		secretAT:                 accessTokenSecret,
		secretRT:                 refreshTokenSecret,
		config:                   c,
	}
}

type Claims struct {
	UserUUID string `json:"userUUID"`
	jwt.StandardClaims
}

// https://developer.vonage.com/blog/20/03/13/using-jwt-for-authentication-in-a-golang-application-dr#:~:text=Refresh%20Token%3A%20A%20refresh%20token,hit%20(from%20our%20application).
func (jt *JWTToken) CreateTokenPair(ctx context.Context, userID string) (*Token, error) {
	ttl := time.Duration(jt.accessTokenExpiryMinutes) * time.Minute
	atExpiryTime := time.Now().Add(ttl)
	rtExpiryTime := time.Now().Add(time.Hour * 24 * 7)

	at, err := jt.NewAccessToken(ctx, &Claims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: atExpiryTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})
	if err != nil {
		return nil, err
	}

	rt, err := jt.NewRefreshToken(ctx, &Claims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: rtExpiryTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})
	if err != nil {
		return nil, err
	}

	return &Token{
		UserUUID:               userID,
		AccessToken:            at,
		RefreshToken:           rt,
		AccessExpiryTime:       atExpiryTime.Unix(),
		RefreshExpiryTime:      rtExpiryTime.Unix(),
		AccessExpiryTimeStamp:  atExpiryTime.UTC().String(),
		RefreshExpiryTimeStamp: rtExpiryTime.UTC().String(),
	}, nil
}

func (jt *JWTToken) NewAccessToken(ctx context.Context, claims *Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(jt.secretAT))
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func (jt *JWTToken) NewRefreshToken(ctx context.Context, claims *Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken, err := token.SignedString([]byte(jt.secretRT))
	if err != nil {
		return "", err
	}
	return refreshToken, nil
}

func (jt *JWTToken) Refresh(ctx context.Context, refreshToken string) (*Token, error) {
	token, err := jwt.ParseWithClaims(refreshToken, &Claims{}, func(token *jwt.Token) (any, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jt.secretRT), nil
	})

	//if there is an error, the token must have expired
	if err != nil {
		return nil, err
	}

	tokenClaims, ok := token.Claims.(*Claims)
	switch {
	case !ok && !token.Valid:
		return nil, err

	case tokenClaims == nil:
		return nil, errors.New("got nil token claims")
	}

	newTokenPair, err := jt.CreateTokenPair(ctx, tokenClaims.UserUUID)
	if err != nil {
		return nil, err
	}

	return newTokenPair, nil
}
