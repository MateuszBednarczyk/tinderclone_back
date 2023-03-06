package middlewares

import (
	"github.com/labstack/echo/v4"

	"tinderclone_back/src/pkg/handlers"
	"tinderclone_back/src/pkg/services"
)

func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		serviceResponse := services.JwtService().IsTokenValid(authHeader)

		if len(serviceResponse.Content) == 0 {
			return c.JSON(serviceResponse.Code, handlers.CreateHandlerResponse(serviceResponse))
		}

		if serviceResponse.Content[0] == false {
			return c.JSON(serviceResponse.Code, handlers.CreateHandlerResponse(serviceResponse))
		}

		if serviceResponse.Content[1].(*services.JwtClaims).IsAdmin == true {
			return next(c)
		}

		return c.JSON(serviceResponse.Code, handlers.CreateHandlerResponse(serviceResponse))
	}
}
