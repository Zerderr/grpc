package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"homework-5/internal/config"
)

func NewDB(ctx context.Context, cfg *config.Config) (*Database, error) {
	dsn := generateDsn(cfg)
	pool, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return newDataBase(pool), nil
}

func generateDsn(cfg *config.Config) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresDbHost, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDb)
}
