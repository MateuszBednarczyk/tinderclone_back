package handlers

import "github.com/labstack/echo/v4"

type IHealthCheckHandler interface {
	HandleHealthCheck(c echo.Context) error
}

type healthCheckHandler struct {
}

func NewHealthCheckHandler() *healthCheckHandler {
	return &healthCheckHandler{}
}

func (h *healthCheckHandler) HandleHealthCheck(c echo.Context) error {
	return c.JSON(200, "Server instance running")
}
