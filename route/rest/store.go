package rest

import (
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/labstack/echo/v4"
)

func (r *Rest) registerStoreService(e *echo.Echo) {
	httpProxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   os.Getenv("HTTP_STORE_HOST"),
	})

	api := e.Group("/api/v1/stores")
	doc := e.Group("/docs/v1/stores")

	api.POST("", echo.WrapHandler(httpProxy))
	api.GET("/:id", echo.WrapHandler(httpProxy))
	api.GET("", echo.WrapHandler(httpProxy))

	doc.GET("", echo.WrapHandler(httpProxy))
	doc.GET("/openapi.yaml", echo.WrapHandler(httpProxy))
}
