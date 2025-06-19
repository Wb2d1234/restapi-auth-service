package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"auth-service/internal/config"
)

type Postgres struct {
	DB *sqlx.DB
}

func NewPostgresStorage(cfg *config.Config) (*Postgres, error) {
	db, err := sqlx.Connect("postgres", cfg.DB_DSN)
	if err != nil {
		return nil, fmt.Errorf("connection error: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping error: %w", err)
	}

	return &Postgres{DB: db}, nil
}
