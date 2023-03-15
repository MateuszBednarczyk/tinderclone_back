package handlers

import (
	"github.com/labstack/echo/v4"

	"tinderclone_back/src/pkg/services"
)

type ICountrierHandler interface {
	HandleSaveCountry(c echo.Context) error
}

type countrierHandler struct {
	countrier services.ICountrier
}

func NewCountrierHandler(countrier services.ICountrier) *countrierHandler {
	return &countrierHandler{
		countrier: countrier,
	}
}

func (h *countrierHandler) HandleSaveCountry(c echo.Context) error {
	countryName := c.QueryParam("name")
	serviceResult := h.countrier.SaveNewCountry(countryName)

	return c.JSON(serviceResult.Code, CreateHandlerResponse(serviceResult))
}
