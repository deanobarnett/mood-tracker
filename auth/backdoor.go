package auth

import (
	"fmt"
	"net/http"

	"github.com/deanobarnett/mood-tracker/config"
	"github.com/labstack/echo/v4"
)

func BackDoor(cfg *config.Config, s *Service) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			if isAllowed(cfg.Env) {
				email := c.QueryParam("as")
				ctx := c.Request().Context()
				user, _ := s.GetUser(ctx, email)
				setRememberToken(c.Request(), user.RememberToken)
			}

			return next(c)
		}
	}
}

func setRememberToken(r *http.Request, token string) {
	// Set Cookie
	r.AddCookie(NewCookie(token))

	// Set Auth Header
	authHeader := fmt.Sprintf("Bearer %s", token)
	r.Header.Set("Authorization", authHeader)
}

func isAllowed(e string) bool {
	testEnvs := []string{"test", "ci", "dev"}

	for _, env := range testEnvs {
		if env == e {
			return true
		}
	}

	return false
}
