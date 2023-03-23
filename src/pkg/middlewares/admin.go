package middlewares

import (
	"github.com/labstack/echo/v4"

	"tinderclone_back/src/pkg/handlers"
	"tinderclone_back/src/pkg/services"
)

func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		serviceResponse := services.Tokenizer().IsTokenValid(authHeader)

		if serviceResponse.Code != 200 {
			return c.JSON(serviceResponse.Code, handlers.CreateHandlerResponse(serviceResponse))
		}

		if serviceResponse.Content[0].(*services.JwtClaims).User.Role != 2 {
			return c.JSON(403, "You don't have permission")
		}

		return next(c)
	}
}
