package setups

import db_gen "github.com/qryne/api/internal/db/sqlc"

type ISetupRepository interface {
	CreateEntry() (db_gen.Setup, error)
}

type ISetupServices interface {
	InitSetup() error
}
