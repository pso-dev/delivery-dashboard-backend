package main

import (
	"os"

	"github.com/pso-dev/delivery-dashboard/backend/internal/app"
	"github.com/pso-dev/delivery-dashboard/backend/pkg/jlog"
)

func main() {
	logger := jlog.New(os.Stdout, jlog.LevelInfo)
	if err := run(os.Args, logger); err != nil {
		logger.PrintError(err, nil)
	}
}

func run(args []string, logger *jlog.Logger) error {
	cfg := app.Configuration{}

	cfg.ENV = "development"
	// TODO: Move these out from being hard-coded. Only for early development
	cfg.DB.DSN = "postgres://postgres:password@localhost/pso?sslmode=disable"
	cfg.DB.MaxOpenConnections = 25
	cfg.DB.MaxIdleConnections = 25
	cfg.DB.MaxIdleTime = "15m"

	app := app.New(cfg, logger)

	return app.Run()
}
