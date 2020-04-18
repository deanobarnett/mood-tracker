package user

import "github.com/labstack/echo/v4"

func Authorize(s *Service) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cookie, err := c.Cookie("remember_token")
			if err != nil {
				return err
			}
			err = s.Validate(c.Request().Context(), cookie.Value)
			if err != nil {
				return err
			}
			return next(c)
		}
	}
}
