package main

import (
	"log"
	"os"
	"time"

	"github.com/deanobarnett/mood-tracker/entry"
	"github.com/deanobarnett/mood-tracker/user"
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

	// set up user service
	userService := user.NewService(db)
	userServer := user.NewHTTPServer(userService)
	userServer.RouteTo(e, user.Authorize(userService))

	// set up entry service
	entryService := entry.NewService(db)
	entryServer := entry.NewHTTPServer(entryService)
	entryServer.RouteTo(e, user.Authorize(userService))

	return graceful.ListenAndServe(e.Server, 5*time.Second)
}

func main() {
	log.Fatal(run())
}
