package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (r *Rest) getUsers(e echo.Context) error {
	return e.JSON(http.StatusOK, echo.Map{"message": "OK"})
}
