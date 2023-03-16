package handlers

import (
	"github.com/labstack/echo/v4"

	"tinderclone_back/src/pkg/services"
)

type IPermitterHandler interface {
	HandleGiveUserAdminPermission(c echo.Context) error
}

type permitterHandler struct {
	permitter services.IPermitter
}

func NewPermitterHandler(permitter services.IPermitter) *permitterHandler {
	return &permitterHandler{
		permitter: permitter,
	}
}

func (h *permitterHandler) HandleGiveUserAdminPermission(c echo.Context) error {
	username := c.QueryParam("username")
	serviceResult := h.permitter.GiveUserAdminPermission(username)

	return c.JSON(serviceResult.Code, CreateHandlerResponse(serviceResult))
}
