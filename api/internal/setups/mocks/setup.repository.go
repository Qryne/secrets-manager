package mocks

import (
	db_gen "github.com/qryne/api/internal/db/sqlc"
	"github.com/stretchr/testify/mock"
)

type ISetupRepository struct {
	mock.Mock
}

func (m *ISetupRepository) CreateEntry() (db_gen.Setup, error) {

	args := m.Called()

	return args.Get(0).(db_gen.Setup), args.Error(1)
}
