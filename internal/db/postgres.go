package db

import (
	"context"
	"fmt"
	"log"

	"github.com/asm2212/go-gin-starter/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresDB(cfg *config.Config) *pgxpool.Pool {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	return pool
}
