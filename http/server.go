package http

import (
	"time"

	"github.com/deanobarnett/mood-tracker/entry"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/tylerb/graceful"
)

type Server struct {
	Router       *echo.Echo
	EntryService *entry.Service
}

func (s *Server) Run() {
	s.Router.Pre(middleware.RemoveTrailingSlash())
	s.Router.Use(middleware.Logger())
	s.Router.Use(middleware.Recover())
	s.Router.Use(middleware.RequestID())
	s.Router.Use(middleware.Secure())
	s.Router.Use(middleware.Gzip())

	s.Router.GET("/entries/:id", s.GetEntry)
	s.Router.GET("/entries", s.ListEntries)
	s.Router.POST("/entries", s.CreateEntry)

	graceful.ListenAndServe(s.Router.Server, 5*time.Second)
}
