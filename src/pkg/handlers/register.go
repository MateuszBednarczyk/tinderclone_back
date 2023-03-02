package handlers

import (
	"github.com/labstack/echo/v4"

	"tinderclone_back/src/pkg/dto"
	"tinderclone_back/src/pkg/services"
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
	service := services.RegisterService()
	var dto dto.RegisterUser
	err := c.Bind(&dto)
	if err != nil {
		return c.JSON(400, "Couldn't read dto")
	}
	serviceResponse := service.RegisterUser(dto)

	return c.JSON(serviceResponse.Code, CreateResponse(serviceResponse))
}
