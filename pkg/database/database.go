package database

import (
	"database/sql"
	"fmt"

	"github.com/begenov/tsarka-task/internal/config"

	_ "github.com/lib/pq"
)

func Open(cfg config.PostgresConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
