package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/qryne/api/internal/db"
	"github.com/qryne/api/internal/env"
)

func main() {
	cfg := Config{
		addr: ":8080",
		db: DBConfig{
			dsn: env.GetString("DB_CONNECTION_STRING", "postgres://qryne:qryne@localhost:5432/qryne?sslmode=disable"),
		},
	}

	// Database
	db := db.InitializeDatabase(cfg.db.dsn)
	ctx := context.Background()
	defer db.Close(ctx)

	// Logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	api := App{
		config: cfg,
		db:     db,
	}

	if err := api.run(api.mount()); err != nil {
		slog.Error("Server failed to start, err: %s", err)
		os.Exit(1)
	}
}
