package handlers

import (
	"github.com/labstack/echo/v4"
)

type IRegisterHandler interface {
	HandleRegister()
}

type registerHandler struct {
}

func NewRegisterHandler() *registerHandler {
	return &registerHandler{}
}

func (h *registerHandler) HandleRegister(c echo.Context) error {
	return c.JSON(200, "hi")
}
