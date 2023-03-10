package handlers

import (
	"github.com/labstack/echo/v4"

	"tinderclone_back/src/pkg/services"
)

type ICountrierHandler interface {
	HandleSaveCountry(c echo.Context) error
}

type countrierHandler struct {
}

func NewCountrierHandler() *countrierHandler {
	return &countrierHandler{}
}

func (h *countrierHandler) HandleSaveCountry(c echo.Context) error {
	countryName := c.QueryParam("name")
	serviceResult := services.Countrier().SaveNewCountry(countryName)

	return c.JSON(serviceResult.Code, CreateHandlerResponse(serviceResult))
}
