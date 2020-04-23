package main

import (
	"log"
	"time"

	"github.com/deanobarnett/mood-tracker/auth"
	"github.com/deanobarnett/mood-tracker/config"
	"github.com/deanobarnett/mood-tracker/entry"
	"github.com/tylerb/graceful"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func run() error {
	config := config.NewConfig()
	e := echo.New()
	e.Server.Addr = ":" + string(config.Port)

	db, err := sqlx.Connect("postgres", config.DB)
	if err != nil {
		return err
	}
	defer db.Close()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.RequestID())
	e.Use(middleware.Secure())
	e.Use(middleware.Gzip())
	e.Use(middleware.CSRF())

	// set up auth service
	authService := auth.NewService(db)
	authServer := auth.NewHTTPServer(authService)
	authServer.RouteTo(e, auth.Authorize(authService))
	// set up auth backdoor
	e.Use(auth.BackDoor(config, authService))

	// set up entry service
	entryService := entry.NewService(db)
	entryServer := entry.NewHTTPServer(entryService)
	entryServer.RouteTo(e, auth.Authorize(authService))

	return graceful.ListenAndServe(e.Server, 5*time.Second)
}

func main() {
	log.Fatal(run())
}
