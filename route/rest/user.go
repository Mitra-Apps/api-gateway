package rest

import (
	"net/http"

	"github.com/Mitra-Apps/be-api-gateway/auth"
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

func (r *Rest) login(e echo.Context) error {
	u := new(pb.UserLoginRequest)
	if err := e.Bind(u); err != nil {
		echo.NewHTTPError(http.StatusBadRequest, err)
	}
	ctx := e.Request().Context()
	user, err := r.userService.Login(ctx, u)
	if err != nil {
		return convertGrpcToHttpErrorResponse(err)
	}
	jwt, err := auth.GenerateToken(ctx, user.User.Username)
	if err != nil {
		echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	response := map[string]interface{}{
		"jwt": jwt,
	}

	return e.JSON(http.StatusOK, response)
}
