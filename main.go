package main

import (
	"flag"
	"log"
	"net/http"

	storePb "github.com/Mitra-Apps/be-api-gateway/domain/proto/store"
	userPb "github.com/Mitra-Apps/be-api-gateway/domain/proto/user"
	"github.com/Mitra-Apps/be-api-gateway/lib"
	"github.com/Mitra-Apps/be-api-gateway/route/rest"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/joho/godotenv"

	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/Mitra-Apps/be-api-gateway/docs"
)

// @title Echo Swagger Example API
// @version 1.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	envInit()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAccessControlAllowHeaders, echo.HeaderAuthorization},
		AllowMethods: []string{http.MethodGet, http.MethodOptions, http.MethodPost, http.MethodDelete, http.MethodPut},
	}))

	userGrpcAddr := flag.String("userGrpcAddr", lib.GetEnv("GRPC_USER_HOST"), "User service host")
	storeGrpcAddr := flag.String("storeGrpcAddr", lib.GetEnv("GRPC_STORE_HOST"), "Store service host")

	userGrpcConn, err := grpc.Dial(*userGrpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Cannot connect to user grpc server ", err)
	}
	storeGrpcConn, err := grpc.Dial(*storeGrpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Cannot connect to store grpc server ", err)
	}
	defer func() {
		log.Println("Closing connection ...")
		storeGrpcConn.Close()
		userGrpcConn.Close()
	}()

	userGrpcServiceClient := userPb.NewUserServiceClient(userGrpcConn)
	storeGrpcServiceClient := storePb.NewStoreServiceClient(storeGrpcConn)
	rest.New(userGrpcServiceClient, storeGrpcServiceClient).Register(e)

	// Health check
	e.GET("/", HealthCheck)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(lib.GetEnv("APP_PORT")))
}

func envInit() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Fatalln("No .env file found")
	}
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": "Server is up and running",
	})
}
