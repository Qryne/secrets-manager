package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InitializeDatabase(dsn string) *pgx.Conn {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, dsn)

	if err != nil {
		panic(err)
	}
	return conn
}
