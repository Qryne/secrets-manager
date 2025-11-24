package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	apikeys "github.com/qryne/api/internal/api_keys"
	"github.com/qryne/api/internal/auth"
	"github.com/qryne/api/internal/db"
	"github.com/qryne/api/internal/setups"
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
	controllers := AuthController{AuthServices: &authServices}
	return controllers
}

func (reg *V1RouterRegister) injectAPIKeyControllers() APIKeyController {
	psqlHandler := &db.PSQLHandler{Conn: reg.PGConn}

	apiKeyRepo := apikeys.APIKeyRepo{
		Db: psqlHandler,
	}
	apiKeyServices := apikeys.APIKeyServices{
		APIKeyRepo: &apiKeyRepo,
	}
	controllers := APIKeyController{APIKeysServices: &apiKeyServices}
	return controllers
}

func (reg *V1RouterRegister) injectSetupControllers() SetupController {
	psqlHandler := &db.PSQLHandler{Conn: reg.PGConn}

	setupRepo := setups.SetupRepo{
		Db: psqlHandler,
	}
	setupServices := setups.SetupServices{
		SetupRepo: &setupRepo,
	}
	controllers := SetupController{SetupServices: &setupServices}
	return controllers
}

func (reg *V1RouterRegister) RegisterV1Router(r chi.Router) {

	authControllers := reg.injectAuthControllers()
	apiKeyContollers := reg.injectAPIKeyControllers()
	setupControlles := reg.injectSetupControllers()

	r.Route("/users", func(r chi.Router) {
		r.Post("/signup", authControllers.UserSignup)
	})
	r.Route("/api-keys", func(r chi.Router) {
		r.Post("/", apiKeyContollers.CreateAPIKeyController)
	})
	r.Route("/setups", func(r chi.Router) {
		r.Post("/", setupControlles.InitSetupController)
	})
}
