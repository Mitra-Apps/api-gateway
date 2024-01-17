package rest

import (
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/labstack/echo/v4"
)

// Get Stores
// @Summary Display all stores
// @Description Display all stores.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} pb.GetStoresResponse
// @Router /api/v1/store [get]
// func (r *Rest) getStores(e echo.Context) error {
// 	StoreList, err := r.storeService.GetStores(e.Request().Context(), &pb.GetStoresRequest{})
// 	if err != nil {
// 		echo.NewHTTPError(http.StatusInternalServerError, err)
// 	}
// 	return e.JSON(http.StatusOK, StoreList)
// }

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
