package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/Mitra-Apps/be-api-gateway/lib"
	jwt "github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type JwtCustomClaim struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(c context.Context, username string) (string, *echo.HTTPError) {
	expireTime, err := time.ParseDuration(lib.GetEnv("JWT_EXPIRED_TIME"))
	if err != nil {
		return "", echo.NewHTTPError(http.StatusBadRequest, "Invalid JWT expired time")
	}
	claims := &JwtCustomClaim{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireTime)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(lib.GetEnv("JWT_SECRET")))
	if err != nil {
		return "", echo.NewHTTPError(http.StatusInternalServerError, "Error when generating token")
	}
	return t, nil
}

func Required() echo.MiddlewareFunc {
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(JwtCustomClaim)
		},
		SigningKey: []byte(lib.GetEnv("JWT_SECRET")),
	}
	return echojwt.WithConfig(config)
}
