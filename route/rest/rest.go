package rest

import (
	"net/http"

	userPb "github.com/Mitra-Apps/be-api-gateway/domain/user/pb"

	"github.com/labstack/echo/v4"
)

type Rest struct {
	userService userPb.UserServiceClient
}

func New(userService userPb.UserServiceClient) *Rest {
	return &Rest{userService}
}

func (r *Rest) Register(e *echo.Echo) {
	e.GET("/ping", r.ping)
	user := e.Group("/api/v1/user")
	user.GET("", r.getUsers)
}

func (r *Rest) ping(e echo.Context) error {
	return e.String(http.StatusOK, "pong")
}
