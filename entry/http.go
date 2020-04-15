package entry

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HTTPServer struct {
	Service *Service
}

func (h *HTTPServer) RouteTo(e *echo.Echo) {
	e.GET("/entries/:id", h.getEntry)
	e.GET("/entries", h.listEntries)
	e.POST("/entries", h.createEntry)
}

func (h *HTTPServer) getEntry(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	entry, err := h.Service.GetEntry(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, entry)
}

func (h *HTTPServer) listEntries(c echo.Context) error {
	perPageParam := c.QueryParam("per_page")
	if perPageParam == "" {
		perPageParam = "10"
	}
	perPage, _ := strconv.ParseInt(perPageParam, 10, 64)
	entries, err := h.Service.ListEntries(perPage)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, entries)
}

func (h *HTTPServer) createEntry(c echo.Context) error {
	entry := &Model{}
	c.Bind(entry)
	entry, err := h.Service.CreateEntry(entry)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, entry)
}
