package auth

import (
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

	c.Request().AddCookie(NewCookie(user.RememberToken))
	return c.JSON(http.StatusCreated, user)
}

// reset password
func (h *HTTPServer) createPassword(c echo.Context) error {
	token := c.QueryParam("remember_token")
	ctx := c.Request().Context()

	err := h.Service.ResetPassword(ctx, token)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, nil)
}

func (h *HTTPServer) updatePassword(c echo.Context) error {
	token := c.QueryParam("confirmation_token")
	id := c.Param("id")
	password := c.FormValue("password")
	ctx := c.Request().Context()

	err := h.Service.UpdatePassword(ctx, id, token, password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, nil)
}

func (h *HTTPServer) createSession(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	ctx := c.Request().Context()

	user, err := h.Service.Login(ctx, email, password)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid credentials")
	}
	c.Request().AddCookie(NewCookie(user.RememberToken))
	return c.JSON(http.StatusCreated, user)
}

// destroy the users current session
func (h *HTTPServer) deleteSession(c echo.Context) error {
	token := c.QueryParam("remember_token")

	err := h.Service.SignOut(c.Request().Context(), token)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid credentials")
	}

	c.Request().AddCookie(ExpireCookie())
	return c.JSON(http.StatusOK, nil)
}
