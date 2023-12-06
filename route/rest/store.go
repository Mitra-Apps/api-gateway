package rest

import (
	"net/http"

	pb "github.com/Mitra-Apps/be-api-gateway/domain/proto/store"
	"github.com/labstack/echo/v4"
)

func (r *Rest) getStores(e echo.Context) error {
	StoreList, err := r.storeService.GetStores(e.Request().Context(), &pb.GetStoresRequest{})
	if err != nil {
		echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return e.JSON(http.StatusOK, StoreList)
}
