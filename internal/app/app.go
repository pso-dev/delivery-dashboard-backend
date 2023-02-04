package app

import (
	"database/sql"
	"fmt"

	"github.com/pso-dev/delivery-dashboard/backend/internal/data"
)

type Configuration struct {
	ENV string
	DB  struct {
		DSN                string
		MaxOpenConnections int
		MaxIdleConnections int
		MaxIdleTime        string
	}
}

type application struct {
	cfg    Configuration
	db     *sql.DB
	models *data.Models
}

func New(cfg Configuration) *application {
	return &application{cfg: cfg}
}

func (a *application) Run() error {
	fmt.Println("Hello World")
	return nil
}
