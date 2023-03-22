package handlers

import (
	"strings"

	"github.com/labstack/echo/v4"

	"tinderclone_back/src/pkg/services"
)

type ICountrierHandler interface {
	HandleSaveCountry(c echo.Context) error
	GetAllCountriesNames(c echo.Context) error
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
	if len(strings.Trim(countryName, "")) == 0 {
		return c.JSON(400, map[string]string{
			"Message": "Given arguments cannot be null or blank",
		})
	}
	serviceResult := h.countrier.SaveNewCountry(countryName)

	return c.JSON(serviceResult.Code, CreateHandlerResponse(serviceResult))
}

func (h *countrierHandler) HandleGetAllCountriesNames(c echo.Context) error {
	serviceResult := h.countrier.GetAllCountriesNames()
	return c.JSON(serviceResult.Code, CreateHandlerResponse(serviceResult))
}
