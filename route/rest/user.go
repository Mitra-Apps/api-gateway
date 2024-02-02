package rest

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	pb "github.com/Mitra-Apps/be-user-service/domain/proto/user"

	"github.com/labstack/echo/v4"
)

// Get Users
// @Summary Display all users
// @Description Display all users.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} pb.GetUsersResponse
// @Router /api/v1/user [get]
func (r *Rest) getUsers(e echo.Context) error {
	userList, err := r.userService.GetUsers(e.Request().Context(), &pb.GetUsersRequest{})
	if err != nil {
		echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return e.JSON(http.StatusOK, userList)
}

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
