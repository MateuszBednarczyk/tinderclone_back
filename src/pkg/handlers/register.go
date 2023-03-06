package handlers

import (
	"github.com/go-playground/validator/v10"
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
	var requestBody dto.RegisterUser
	var err error

	err = c.Bind(&requestBody)
	if err != nil {
		return c.JSON(400, "Couldn't read dto")
	}

	v := validator.New()
	err = v.Struct(requestBody)

	if err != nil {
		return c.JSON(400, err.Error())
	}
	serviceResponse := services.RegisterService().RegisterUser(requestBody)

	return c.JSON(serviceResponse.Code, CreateResponse(serviceResponse))
}
