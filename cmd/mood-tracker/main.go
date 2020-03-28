package main

import (
	"github.com/deanobarnett/mood-tracker/entry"
	"github.com/deanobarnett/mood-tracker/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

func main() {
	e := echo.New()
	e.Server.Addr = ":8080"

	db, err := sqlx.Connect("postgres", "host=db user=postgres dbname=postgres password=postgres sslmode=disable")
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer db.Close()

	entryService := &entry.Service{DB: db}
	server := &http.Server{Router: e, EntryService: entryService}

	server.Run()
}
