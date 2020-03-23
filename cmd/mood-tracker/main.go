package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/deanobarnett/mood-tracker/handler"
)

func main() {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.Secure())

	e.GET("/entries", handler.GetEntries)
	e.POST("/entries", handler.CreateEntry)

	e.Logger.Fatal(e.Start(":8080"))
}
