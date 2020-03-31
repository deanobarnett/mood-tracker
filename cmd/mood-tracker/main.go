package main

import (
	"os"

	"github.com/deanobarnett/mood-tracker/entry"
	"github.com/deanobarnett/mood-tracker/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	e := echo.New()
	e.Server.Addr = ":8080"

	dbURL := os.Getenv("DATABASE_URL")
	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer db.Close()

	entryService := &entry.Service{DB: db}
	server := &http.Server{Router: e, EntryService: entryService}

	server.Run()
}
