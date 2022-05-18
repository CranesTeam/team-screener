package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func NewPostgresDb(ctx context.Context, cfg Config) (*pgx.Conn, error) {
	connectionUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	return pgx.Connect(ctx, connectionUrl)
}
