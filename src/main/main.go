package main

import (
	"fmt"
	"sync"

	"github.com/labstack/echo/v4"
)

var serverInstance *echo.Echo

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan string)
	wg.Add(1)
	go launchServer(&wg, ch)
	fmt.Println(<-ch)
	wg.Wait()
}
