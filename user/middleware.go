package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Authorize(s *Service) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.QueryParam("remember_token")
			if token == "" {
				cookie, err := c.Cookie("remember_token")
				if err != nil {
					return echo.NewHTTPError(http.StatusUnauthorized, "missing auth token")
				}
				token = cookie.Value
			}

			err := s.Validate(c.Request().Context(), token)
			if err != nil {
				c.Logger().Warnf("unauthorized user access: %s", err.Error())
				return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized user")
			}

			return next(c)
		}
	}
}
