package app

import (
	"database/sql"
	"sync"

	"github.com/pso-dev/delivery-dashboard/backend/internal/data"
	"github.com/pso-dev/delivery-dashboard/backend/internal/data/postgres"
	"github.com/pso-dev/delivery-dashboard/backend/pkg/jlog"
)

type Configuration struct {
	Port int64
	ENV  string
	DB   struct {
		DSN                string
		MaxOpenConnections int
		MaxIdleConnections int
		MaxIdleTime        string
	}
}

type application struct {
	cfg          Configuration
	logger       *jlog.Logger
	db           *sql.DB
	repositories *data.Repositories
	wg           sync.WaitGroup
	mu           sync.Mutex
}

func New(cfg Configuration, logger *jlog.Logger) *application {
	return &application{
		cfg:    cfg,
		logger: logger,
		mu:     sync.Mutex{},
	}
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

	a.logger.PrintInfo("DB connection pool established", map[string]string{"connectionPoolSize": "25"})

	a.db = db
	a.repositories = data.NewRepositories(db)

	return a.serve()
}
