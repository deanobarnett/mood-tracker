package handler

import (
	"net/http"
	"time"

	"github.com/deanobarnett/mood-tracker/model"
	"github.com/labstack/echo"
)

func GetEntries(c echo.Context) error {
	entry := &model.Entry{
		ID:   1,
		Date: time.Now().Format(time.RFC3339),
	}
	return c.JSON(http.StatusOK, entry)
}

func CreateEntry(c echo.Context) error {
	entry := new(model.Entry)
	if err := c.Bind(entry); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, entry)
}
