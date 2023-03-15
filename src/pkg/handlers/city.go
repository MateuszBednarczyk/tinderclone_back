package handlers

import (
	"github.com/labstack/echo/v4"

	"tinderclone_back/src/pkg/services"
)

type ICitierHandler interface {
	HandleSaveNewCity(c echo.Context) error
}

type citierHandler struct {
	citier services.ICitier
}

func NewCityHandler(citier services.ICitier) *citierHandler {
	return &citierHandler{
		citier: citier,
	}
}

func (h *citierHandler) HandleSaveNewCity(c echo.Context) error {
	cityName := c.QueryParam("name")
	countryName := c.QueryParam("country")
	serviceResult := h.citier.SaveNewCity(cityName, countryName)

	return c.JSON(serviceResult.Code, CreateHandlerResponse(serviceResult))
}
