package user

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HTTPServer struct {
	Service *Service
}

func NewHTTPServer(svc *Service) *HTTPServer {
	return &HTTPServer{Service: svc}
}

func (h *HTTPServer) RouteTo(e *echo.Echo, auth echo.MiddlewareFunc) {
	e.POST("/users", h.createUser)

	e.POST("/passwords", h.createPassword)
	e.PUT("/users/:id/passwords", h.updatePassword, auth)

	e.POST("/sessions", h.createSession)
	e.DELETE("/sessions", h.deleteSession, auth)
}

func (h *HTTPServer) createUser(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	ctx := c.Request().Context()

	user, err := h.Service.CreateUser(ctx, email, password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, user)
}

func (h *HTTPServer) createPassword(c echo.Context) error {
	return fmt.Errorf("not implemented")
}

func (h *HTTPServer) updatePassword(c echo.Context) error {
	return fmt.Errorf("not implemented")
}

func (h *HTTPServer) createSession(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	ctx := c.Request().Context()

	user, err := h.Service.Login(ctx, email, password)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid credentials")
	}
	return c.JSON(http.StatusCreated, user)
}

func (h *HTTPServer) deleteSession(c echo.Context) error {
	return fmt.Errorf("not implemented")
}
