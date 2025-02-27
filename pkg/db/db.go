package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rotisserie/eris"

	"github.com/sashajdn/replog/pkg/env"
)

type DB struct {
	*pgxpool.Pool
}

func New(cfg env.PostgreSQL) (*DB, error) {
	dbConfig, err := pgxpool.ParseConfig(cfg.URL)
	if err != nil {
		return nil, eris.Wrap(err, "parse db config")
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		return nil, eris.Wrap(err, `connect db pool`)
	}

	return &DB{pool}, nil
}
