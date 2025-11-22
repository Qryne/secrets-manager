package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/qryne/api/internal/auth"
	"github.com/qryne/api/internal/db"
)

type V1RouterRegister struct {
	PGConn *pgx.Conn
}

func (reg *V1RouterRegister) injectAuthControllers() AuthController {
	psqlHandler := &db.PSQLHandler{Conn: reg.PGConn}

	authRepo := auth.AuthRepo{
		DBHandler: psqlHandler,
	}
	authServices := auth.AuthService{
		AuthRepo: &auth.AuthRepoWithCktBrkr{AuthRepo: authRepo},
	}
	authController := AuthController{AuthServices: &authServices}
	return authController
}

func (reg *V1RouterRegister) RegisterV1Router(r chi.Router) {

	authControllers := reg.injectAuthControllers()

	r.Route("/users", func(r chi.Router) {
		r.Post("/signup", authControllers.UserSignup)
	})
}
