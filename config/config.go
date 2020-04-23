package config

import (
	"flag"
	"os"

	"github.com/peterbourgon/ff"
)

type Config struct {
	DB    string
	Env   string
	Port  int
	Debug bool
}

func NewConfig() *Config {
	fs := flag.NewFlagSet("mood-tracker", flag.ExitOnError)

	var (
		port  = fs.Int("port", 8080, "listen port for server")
		debug = fs.Bool("debug", false, "log debug information")
		env   = fs.String("env", "dev", "environmenrt")
		db    = fs.String("db_dsn", "", "database url")
	)

	ff.Parse(fs, os.Args[1:], ff.WithEnvVarNoPrefix())

	return &Config{
		DB:    *db,
		Env:   *env,
		Port:  *port,
		Debug: *debug,
	}
}
