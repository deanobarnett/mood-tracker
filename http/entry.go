package http

import (
	"net/http"
	"strconv"
	"time"

	"github.com/deanobarnett/mood-tracker/entry"
	"github.com/labstack/echo"
)

func (s *Server) GetEntry(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	entry, err := s.EntryService.GetEntry(id)
	if err != nil {
		res := &errorResponse{id, err.Error()}
		return c.JSON(http.StatusBadRequest, res)
	}
	return c.JSON(http.StatusOK, entry)
}

func (s *Server) ListEntries(c echo.Context) error {
	entries, err := s.EntryService.ListEntries()
	if err != nil {
		res := &errorResponse{0, err.Error()}
		return c.JSON(http.StatusBadRequest, res)
	}
	return c.JSON(http.StatusOK, entries)
}

func (s *Server) CreateEntry(c echo.Context) error {
	entry := &entry.Model{}
	c.Bind(entry)
	entry, err := s.EntryService.CreateEntry(entry)
	if err != nil {
		res := &errorResponse{0, err.Error()}
		return c.JSON(http.StatusBadRequest, res)
	}
	return c.JSON(http.StatusCreated, entry)
}

type errorResponse struct {
	ID    int64  `json:"id" `
	Error string `json:"error"`
}

type response struct {
	ID     int64     `json:"id"`
	Date   time.Time `json:"date"`
	Mood   int       `json:"mood"`
	Sleep  int       `json:"sleep"`
	Stress int       `json:"stress"`
}
