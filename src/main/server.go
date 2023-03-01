package main

import (
	"sync"

	"github.com/labstack/echo/v4"

	"tinderclone_back/src/pkg/handlers"
	"tinderclone_back/src/pkg/services"
)

func launchServer(wg *sync.WaitGroup, serverInstance *echo.Echo, ch chan string) {
	if serverInstance != nil {
		ch <- "Server instance already running"
		wg.Done()
	}
	serverInstance = echo.New()
	initializeHandlers(serverInstance)
	services.InitializeServices()
	serverInstance.Logger.Fatal(serverInstance.Start("localhost:8000"))
	wg.Done()
}

func initializeHandlers(serverInstance *echo.Echo) {
	healthCheckHandler := handlers.NewHealthCheckHandler()
	serverInstance.GET("/api/v1/health", healthCheckHandler.HandleHealthCheck)
}
