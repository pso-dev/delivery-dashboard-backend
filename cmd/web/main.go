package main

import (
	"fmt"
	"os"

	"github.com/pso-dev/delivery-dashboard/backend/internal/app"
)

func main() {
	if err := run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s/n", err)
	}
}

func run(args []string) error {
	cfg := app.Configuration{}

	cfg.ENV = "development"
	cfg.DB.DSN = ""
	cfg.DB.MaxOpenConnections = 25
	cfg.DB.MaxIdleConnections = 25
	cfg.DB.MaxIdleTime = "15m"

	app := app.New(cfg)

	return app.Run()
}
