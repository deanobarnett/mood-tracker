package main

import (
	"os"
	"time"

	"github.com/deanobarnett/mood-tracker/entry"
	"github.com/tylerb/graceful"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func main() {
	e := echo.New()
	e.Server.Addr = ":8080"

	dbURL := os.Getenv("DB_DSN")
	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer db.Close()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.Secure())
	e.Use(middleware.Gzip())

	{
		entryService := &entry.Service{DB: db}
		entryServer := &entry.HTTPServer{Service: entryService}
		entryServer.RouteTo(e)
	}

	graceful.ListenAndServe(e.Server, 5*time.Second)
}
