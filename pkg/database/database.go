package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sumukhj1219/go-profiler/internals/config"
	"github.com/sumukhj1219/go-profiler/pkg/logger"
)

func New(ctx context.Context, cfg config.Config) (db *sqlx.DB, err error) {
	dsn := cfg.DatabaseConfig.GetDSN()
	logger.L.Info("Database is connecting...")
	db, err = sqlx.Open("postgres", dsn)
	if err != nil {
		fmt.Printf("Unable to connect database %v", err)
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := db.PingContext(pingCtx); err != nil {
		return nil, err
	}

	logger.L.Info("Database connected successfully")
	return db, nil
}
