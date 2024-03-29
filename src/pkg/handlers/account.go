package handlers

import (
	"github.com/labstack/echo/v4"

	"tinderclone_back/src/pkg/services"
)

type IAccounterHandler interface {
	GetAccountInformations(c echo.Context) error
}

type accounterHandler struct {
	accounter services.IAccounter
}

func NewAccounterHandler(accounter services.IAccounter) *accounterHandler {
	return &accounterHandler{
		accounter: accounter,
	}
}

func (h *accounterHandler) GetAccountInformations(c echo.Context) error {
	username := c.QueryParam("username")
	serviceResult := h.accounter.GetAccountInformations(username)

	return c.JSON(serviceResult.Code, CreateHandlerResponse(serviceResult))
}
