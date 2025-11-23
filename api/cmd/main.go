package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/qryne/api/internal/db"
	"github.com/qryne/api/utility"
)

func main() {

	dsn, err := utility.GetString("DB_CONNECTION_STRING")
	if err != nil {
		slog.Error("Server failed to start, err: %s", err)
		os.Exit(1)
	}
	cfg := Config{
		addr: ":8080",
		db: DBConfig{
			dsn,
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
