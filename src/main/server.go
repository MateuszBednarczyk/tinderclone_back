package main

import (
	"sync"

	"github.com/labstack/echo/v4"
)

func launchServer(wg *sync.WaitGroup, serverInstance *echo.Echo, ch chan string) {
	if serverInstance != nil {
		ch <- "Server instance already running"
		wg.Done()
	}
	serverInstance = echo.New()
	serverInstance.Logger.Fatal(serverInstance.Start("localhost:8000"))
	wg.Done()
}
