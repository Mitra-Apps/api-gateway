package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Rest struct {
}

func New() *Rest {
	return &Rest{}
}

func (r *Rest) Register(e *echo.Echo) {
	e.GET("/ping", r.ping)
	user := e.Group("/api/v1/user")
	user.GET("", r.getUsers)
}

func (r *Rest) ping(e echo.Context) error {
	return e.String(http.StatusOK, "pong")
}
