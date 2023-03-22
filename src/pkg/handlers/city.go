package handlers

import (
	"strings"

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
	if len(strings.Trim(cityName, "")) == 0 || len(strings.Trim(countryName, "")) == 0 {
		return c.JSON(400, map[string]string{
			"Message": "Given arguments cannot be null or blank",
		})
	}

	serviceResult := h.citier.SaveNewCity(cityName, countryName)

	return c.JSON(serviceResult.Code, CreateHandlerResponse(serviceResult))
}
