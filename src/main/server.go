package main

import (
	"sync"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"

	"tinderclone_back/src/pkg/database"
	"tinderclone_back/src/pkg/handlers"
	"tinderclone_back/src/pkg/middlewares"
	"tinderclone_back/src/pkg/services"
)

var (
	server     = "localhost"
	port       = "8000"
	apiVersion = "v1"
	dbUsername = "root"
	dbPassword = ""
	dbPort     = "5432"
	dbHost     = "localhost"
	dbName     = "tinder"
)

func launchServer(wg *sync.WaitGroup, ch chan string) {
	if serverInstance != nil {
		ch <- "Server instance already running"
		wg.Done()
	}
	serverInstance = echo.New()
	database.InitializeDb(&database.DbConfig{
		DbUsername: dbUsername,
		DbPassword: dbPassword,
		DbPort:     dbPort,
		DbHost:     dbHost,
		DbName:     dbName,
	})
	initializeHandlers(serverInstance)
	services.InitializeServices()
	serverInstance.Logger.Fatal(serverInstance.Start(server + ":" + port))
	wg.Done()
}

func initializeHandlers(si *echo.Echo) {
	healthCheckHandler := handlers.NewHealthCheckHandler()
	registerHandler := handlers.NewRegisterHandler()
	loginHandler := handlers.NewLoginHandler()
	accountHandler := handlers.NewAccounterHandler()
	countrierHandler := handlers.NewCountrierHandler()

	adminGroup := serverInstance.Group("api/" + apiVersion + "/health")
	adminGroup.Use(echojwt.JWT([]byte("secret")))
	si.GET("api/"+apiVersion+"/health", healthCheckHandler.HandleHealthCheck, middlewares.AdminMiddleware)

	userGroup := serverInstance.Group("/api/" + apiVersion + "/user")
	userGroup.Use(echojwt.JWT([]byte("secret")))
	si.GET("api/"+apiVersion+"/user/:username", accountHandler.GetAccountInformations, middlewares.LoggedUserMiddleware)

	si.POST("api/"+apiVersion+"/user", registerHandler.HandleRegister)
	si.POST("api/"+apiVersion+"/auth", loginHandler.HandleLogin)

	si.POST("api/"+apiVersion+"/country", countrierHandler.HandleSaveCountry)
}
