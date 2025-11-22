package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	v1 "github.com/qryne/api/cmd/router/v1"
)

type RegisterRouter struct {
	PGConn *pgx.Conn
}

func (reg *RegisterRouter) RegisterCombinedRouter(r chi.Router) {
	v1Routes := v1.V1RouterRegister{
		PGConn: reg.PGConn,
	}

	r.Route("/v1", func(r chi.Router) { v1Routes.RegisterV1Router(r) })
}
