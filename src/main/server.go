package main

import (
	"sync"

	"github.com/labstack/echo/v4"

	"tinderclone_back/src/pkg/database"
	"tinderclone_back/src/pkg/handlers"
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

func launchServer(wg *sync.WaitGroup, serverInstance *echo.Echo, ch chan string) {
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
	serverInstance.Logger.Fatal(serverInstance.Start("localhost:8000"))
	wg.Done()
}

func initializeHandlers(serverInstance *echo.Echo) {
	healthCheckHandler := handlers.NewHealthCheckHandler()
	registerHandler := handlers.NewRegisterHandler()
	loginHandler := handlers.NewLoginHandler()

	serverInstance.GET("api/"+apiVersion+"/health", healthCheckHandler.HandleHealthCheck)
	serverInstance.POST("api/"+apiVersion+"/user", registerHandler.HandleRegister)
	serverInstance.POST("api/"+apiVersion+"/auth", loginHandler.HandleLogin)
}