package app

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/pso-dev/delivery-dashboard/backend/internal/data"
	"github.com/pso-dev/delivery-dashboard/backend/internal/data/postgres"
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
	cfg          Configuration
	db           *sql.DB
	repositories *data.Repositories
	mu           sync.Mutex
}

func New(cfg Configuration) *application {
	return &application{cfg: cfg}
}

func (a *application) Run() error {

	db, err := postgres.OpenDB(
		a.cfg.DB.DSN,
		a.cfg.DB.MaxOpenConnections,
		a.cfg.DB.MaxIdleConnections,
		a.cfg.DB.MaxIdleTime)
	if err != nil {
		return err
	}
	defer db.Close()

	fmt.Println("DB connection pool established")

	a.db = db
	a.repositories = data.NewRepositories(db)

	a.mu = sync.Mutex{}

	return nil
}
