package main

import (
	"flag"
	"log"
	"net/http"

	userPb "github.com/Mitra-Apps/be-api-gateway/domain/user/proto"
	"github.com/Mitra-Apps/be-api-gateway/lib"
	"github.com/Mitra-Apps/be-api-gateway/route/rest"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/joho/godotenv"
)

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

	userGrpcAddr := flag.String("addr", lib.GetEnv("GRPC_USER_HOST"), "User service host")
	userGrpcConn, err := grpc.Dial(*userGrpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Cannot connect to grpc server ", err)
	}
	defer func() {
		log.Println("Closing connection ...")
		userGrpcConn.Close()
	}()

	userGrpCServiceClient := userPb.NewUserServiceClient(userGrpcConn)
	rest.New(userGrpCServiceClient).Register(e)

	e.Logger.Fatal(e.Start(lib.GetEnv("APP_PORT")))
}

func envInit() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Fatalln("No .env file found")
	}
}
