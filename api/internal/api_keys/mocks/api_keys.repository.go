package mocks

import (
	db_gen "github.com/qryne/api/internal/db/sqlc"
	"github.com/stretchr/testify/mock"
)

type IAPIKeyRepository struct {
	mock.Mock
}

func (m *IAPIKeyRepository) CreateAPIKey(
	name, slug, prefix, publicID, encryptionIV, encryptedText, algorithm, setupID string,
	scope []string,
) (db_gen.ApiKey, error) {

	args := m.Called(name, slug, prefix, publicID, encryptionIV, encryptedText, algorithm, setupID, scope)

	return args.Get(0).(db_gen.ApiKey), args.Error(1)
}
