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

	prodApi := e.Group("/api/v1/products")
	prodCatApi := e.Group("/api/v1/product-category")
	prodTypeApi := e.Group("/api/v1/product-type")
	uomApi := e.Group("/api/v1/uom")

	api.POST("", echo.WrapHandler(httpProxy))
	api.GET("/:id", echo.WrapHandler(httpProxy))
	api.GET("/my-store", echo.WrapHandler(httpProxy))
	api.GET("", echo.WrapHandler(httpProxy))
	api.PUT("/active-toggle/:is_active", echo.WrapHandler(httpProxy))
	api.PUT("/:id", echo.WrapHandler(httpProxy))
	api.DELETE("/:id", echo.WrapHandler(httpProxy))

	doc.GET("", echo.WrapHandler(httpProxy))
	doc.GET("/openapi.yaml", echo.WrapHandler(httpProxy))

	prodApi.GET("/:product_id", echo.WrapHandler(httpProxy))
	e.GET("/api/v1/product-list/:store_id/:is_include_deactivated", echo.WrapHandler(httpProxy))
	prodApi.POST("", echo.WrapHandler(httpProxy))
	prodCatApi.GET("/:is_include_deactivated", echo.WrapHandler(httpProxy))
	prodCatApi.POST("", echo.WrapHandler(httpProxy))
	prodTypeApi.GET("/:is_include_deactivated", echo.WrapHandler(httpProxy))
	prodTypeApi.POST("", echo.WrapHandler(httpProxy))
	uomApi.GET("/:is_include_deactivated", echo.WrapHandler(httpProxy))
	uomApi.POST("", echo.WrapHandler(httpProxy))
	uomApi.PUT("/:uom_id", echo.WrapHandler(httpProxy))
}
