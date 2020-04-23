package auth

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func Authorize(s *Service) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			err := s.Validate(c.Request().Context(), fetchToken(c))

			if err != nil {
				c.Logger().Warnf("unauthorized user access: %s", err.Error())
				return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized user")
			}

			return next(c)
		}
	}
}

func fetchToken(c echo.Context) string {
	token := ""
	// Precedence
	// 1. Cookie
	// 2. Query Param
	// 3. Header
	cookie, err := c.Cookie("remember_token")

	if err == nil {
		token = cookie.Value
	}

	if token == "" {
		token = c.QueryParam("remember_token")
	}

	if token == "" {
		token = parseHeader(c.Request())
	}

	return token
}

// Parse a header of Format "Bearer <token>"
func parseHeader(r *http.Request) string {
	header := r.Header.Get("Authorization")
	return strings.Split(header, " ")[1]
}
