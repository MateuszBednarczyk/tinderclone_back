package handlers

import (
	"github.com/labstack/echo/v4"

	"tinderclone_back/src/pkg/dto"
	"tinderclone_back/src/pkg/services"
)

type ILoginHandler interface {
	HandleLogin(c echo.Context) error
}

type loginHandler struct {
	authenticator services.IAuthenticator
}

func NewLoginHandler(authenticator services.IAuthenticator) *loginHandler {
	return &loginHandler{
		authenticator: authenticator,
	}
}

func (h *loginHandler) HandleLogin(c echo.Context) error {
	var requestBody dto.Credentials
	err := c.Bind(&requestBody)
	if err != nil {
		return c.JSON(400, "Couldn't read the dto")
	}
	serviceResult := h.authenticator.LoginUser(requestBody)

	return c.JSON(serviceResult.Code, CreateHandlerResponse(serviceResult))
}
