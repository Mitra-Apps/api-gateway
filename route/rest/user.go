package rest

import (
	"net/http"

	pb "github.com/Mitra-Apps/be-api-gateway/domain/proto/user"

	"github.com/labstack/echo/v4"
)

func (r *Rest) getUsers(e echo.Context) error {
	userList, err := r.userService.GetUsers(e.Request().Context(), &pb.GetUsersRequest{})
	if err != nil {
		echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return e.JSON(http.StatusOK, userList)
}
