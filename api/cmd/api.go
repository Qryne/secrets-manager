package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"

	"github.com/qryne/api/cmd/router"
)

type App struct {
	config Config
	db     *pgx.Conn
}

func (app *App) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	registerRouter := router.RegisterRouter{
		PGConn: app.db,
	}

	r.Route("/api", func(r chi.Router) { registerRouter.RegisterCombinedRouter(r) })
	return r
}

func (app *App) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 60,
		IdleTimeout:  time.Minute * 2,
	}

	log.Printf("Server running at %s", app.config.addr)
	return srv.ListenAndServe()
}

type Config struct {
	addr string
	db   DBConfig
}

type DBConfig struct {
	dsn string
}
