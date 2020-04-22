package main

import (
	"log"
	"os"
	"time"

	"github.com/deanobarnett/mood-tracker/auth"
	"github.com/deanobarnett/mood-tracker/entry"
	"github.com/tylerb/graceful"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func run() error {
	e := echo.New()
	e.Server.Addr = ":8080"

	dbURL := os.Getenv("DB_DSN")
	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		return err
	}
	defer db.Close()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.RequestID())
	e.Use(middleware.Secure())
	e.Use(middleware.Gzip())

	// set up auth service
	authService := auth.NewService(db)
	authServer := auth.NewHTTPServer(authService)
	authServer.RouteTo(e, auth.Authorize(authService))

	// set up entry service
	entryService := entry.NewService(db)
	entryServer := entry.NewHTTPServer(entryService)
	entryServer.RouteTo(e, auth.Authorize(authService))

	return graceful.ListenAndServe(e.Server, 5*time.Second)
}

func main() {
	log.Fatal(run())
}
