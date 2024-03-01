package rest

import (
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/labstack/echo/v4"
)

func (r *Rest) registerUserService(e *echo.Echo) {
	httpProxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   os.Getenv("HTTP_USER_HOST"),
	})
	doc := e.Group("/docs/v1/users")
	// doc group api for swagger and open api doc
	doc.GET("", echo.WrapHandler(httpProxy))
	doc.GET("/openapi.yaml", echo.WrapHandler(httpProxy))

	user := e.Group("/api/v1/users")
	// user group api
	user.GET("", echo.WrapHandler(httpProxy))
	user.POST("/login", echo.WrapHandler(httpProxy))
	user.GET("/getrole", echo.WrapHandler(httpProxy))
	user.POST("/register", echo.WrapHandler(httpProxy))
	user.POST("/createrole", echo.WrapHandler(httpProxy))
	user.POST("/verify-token", echo.WrapHandler(httpProxy))
	user.POST("/resend-otp", echo.WrapHandler(httpProxy))
	user.POST("/change-password", echo.WrapHandler(httpProxy))
}
