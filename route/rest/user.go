package rest

import (
	"net/http"

	"github.com/Mitra-Apps/api-gateway/domain/user/pb"

	"github.com/labstack/echo/v4"
)

func (r *Rest) getUsers(e echo.Context) error {
	userList, err := r.userService.GetUsers(e.Request().Context(), &pb.GetUserListRequest{})
	if err != nil {
		echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return e.JSON(http.StatusOK, userList)
}
