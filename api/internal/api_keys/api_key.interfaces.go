package apikeys

import db_gen "github.com/qryne/api/internal/db/sqlc"

type IAPIKeyRepository interface {
	CreateAPIKey(name, slug, prefix, public_id, encryption_iv, encrypted_text, algorithm, setup_id string, scope []string) (db_gen.ApiKey, error)
	FindAPIKeys() error
}

type IAPIKeyServices interface {
	GenerateAPIKey(name string) (db_gen.ApiKey, error)
}
