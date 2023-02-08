package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/pso-dev/delivery-dashboard/backend/internal/app"
	"github.com/pso-dev/delivery-dashboard/backend/pkg/jlog"
)

var Version string
var Buildtime string

func main() {
	logger := jlog.New(os.Stdout, jlog.LevelInfo)
	if err := run(os.Args, logger); err != nil {
		logger.PrintError(err, nil)
	}
}

func run(args []string, logger *jlog.Logger) error {
	cfg := app.Configuration{}
	cfg.Version = fmt.Sprintf("%s - %s", Version, Buildtime)

	flags := flag.NewFlagSet(args[0], flag.ExitOnError)

	// TODO: Move these out from being hard-coded. Only for early development
	cfg.Port = *flags.Int64("port", 6543, "port to listen on")
	cfg.Env = *flags.String("env", "development", "executiion environment (development|production)")
	cfg.DB.DSN = *flags.String("db-dsn", "postgres://postgres:password@localhost/pso?sslmode=disable", "database data source name")
	cfg.DB.MaxOpenConnections = *flags.Int("db-max-open-conns", 25, "database maximum open connections")
	cfg.DB.MaxIdleConnections = *flags.Int("db-max-idle-conns", 25, "database maximum idle connections")
	cfg.DB.MaxIdleTime = *flags.String("db-max-idle-time", "15m", "database maximum idle time")

	flags.Func("cors-trusted-origins", "tructed origins (space separated list)", func(val string) error {
		cfg.CORS.TrustedOrigins = strings.Fields(val)
		return nil
	})

	if err := flags.Parse(args[1:]); err != nil {
		return err
	}

	app := app.New(cfg, logger)

	return app.Run()
}
