package middlewares

import (
	"github.com/labstack/echo/v4"

	"tinderclone_back/src/pkg/handlers"
	"tinderclone_back/src/pkg/services"
)

func LoggedUserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		serviceResponse := services.Tokenizer().IsTokenValid(authHeader)
		if serviceResponse.Code != 200 {
			return c.JSON(serviceResponse.Code, handlers.CreateHandlerResponse(serviceResponse))
		}

		serviceResponse.Content = []interface{}{}
		userDTO, err := services.Tokenizer().GetUserDTOFromToken(authHeader)
		if err != nil {
			return c.JSON(500, "Unexpected error while parsing and user")
		}
		c.Set("loggedUserDTO", userDTO)

		return next(c)
	}
}
