package rest

import (
	"net/http"

	storePb "github.com/Mitra-Apps/be-api-gateway/domain/proto/store"
	userPb "github.com/Mitra-Apps/be-api-gateway/domain/proto/user"

	"github.com/labstack/echo/v4"
)

type Rest struct {
	userService  userPb.UserServiceClient
	storeService storePb.StoreServiceClient
}

func New(userService userPb.UserServiceClient, storeService storePb.StoreServiceClient) *Rest {
	return &Rest{userService, storeService}
}

func (r *Rest) Register(e *echo.Echo) {
	e.GET("/ping", r.ping)
	user := e.Group("/api/v1/user")
	user.GET("", r.getUsers)
	store := e.Group("/api/v1/store")
	store.GET("", r.getStores)
}

func (r *Rest) ping(e echo.Context) error {
	return e.String(http.StatusOK, "pong")
}
