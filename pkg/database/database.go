package database

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/sumukhj1219/go-profiler/internals/config"
	"github.com/sumukhj1219/go-profiler/pkg/logger"
)

func New(ctx context.Context) (db *sqlx.DB, err error) {
	dsn := config.Cfg.GetDSN()
	logger.L.Info("Database is connecting...")
	db, err = sqlx.Open("postgres", dsn)
	if err != nil {
		logger.L.Error("Unable to connect database %v", err)
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
