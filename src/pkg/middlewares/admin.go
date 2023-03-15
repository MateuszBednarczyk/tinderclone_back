package middlewares

import (
	"github.com/labstack/echo/v4"

	"tinderclone_back/src/pkg/services"
)

func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		serviceResponse := services.Tokenizer().IsTokenValid(authHeader)

		if !serviceResponse.Content[0].(bool) {
			return c.JSON(403, "")
		}

		if serviceResponse.Content[1].(*services.JwtClaims).Role != 2 {
			return c.JSON(403, "You don't have permission")
		}

		return next(c)
	}
}
