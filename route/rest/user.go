package rest

import (
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/labstack/echo/v4"
)

func (r *Rest) registerUserService(e *echo.Group) {
	httpProxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   os.Getenv("HTTP_USER_HOST"),
	})

	e.GET("", echo.WrapHandler(httpProxy))
	e.POST("/login", echo.WrapHandler(httpProxy))
	e.GET("/getrole", echo.WrapHandler(httpProxy))
	e.POST("/register", echo.WrapHandler(httpProxy))
	e.POST("/createrole", echo.WrapHandler(httpProxy))
	e.POST("/verify-token", echo.WrapHandler(httpProxy))
	e.POST("/resend-otp", echo.WrapHandler(httpProxy))
}
