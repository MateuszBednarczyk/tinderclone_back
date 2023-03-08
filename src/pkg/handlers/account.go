package handlers

import (
	"github.com/labstack/echo/v4"

	"tinderclone_back/src/pkg/services"
)

type IAccounterHandler interface {
	GetAccountInformations(c echo.Context) error
}

type accounterHandler struct {
}

func NewAccounterHandler() *accounterHandler {
	return &accounterHandler{}
}

func (h *accounterHandler) GetAccountInformations(c echo.Context) error {
	username := c.QueryParam("username")
	serviceResult := services.Accounter().GetAccountInformations(username)

	return c.JSON(serviceResult.Code, CreateHandlerResponse(serviceResult))
}
