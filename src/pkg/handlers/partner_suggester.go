package handlers

import (
	"github.com/labstack/echo/v4"

	"tinderclone_back/src/pkg/dto"
	"tinderclone_back/src/pkg/services"
)

type IPartnerSuggesterHandler interface {
	HandleGetSuggestPartners(c echo.Context) error
}

type partnerSuggesterHandler struct {
	partnerSuggester services.IPartnerSuggester
}

func NewPartnerSuggesterHandler(partnerSuggester services.IPartnerSuggester) *partnerSuggesterHandler {
	return &partnerSuggesterHandler{
		partnerSuggester: partnerSuggester,
	}
}

func (h *partnerSuggesterHandler) HandleGetSuggestPartners(c echo.Context) error {
	userDTO := c.Get("loggedUserDTO").(*dto.User)
	if userDTO == nil {
		return c.JSON(500, map[string]string{
			"Message": "Couldn't read user",
		})
	}
	serviceResult := services.PartnerSuggester().SuggestPartners(*userDTO)

	return c.JSON(serviceResult.Code, CreateHandlerResponse(serviceResult))
}
