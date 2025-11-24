package mocks

import (
	apikeys "github.com/qryne/api/internal/api_keys"
	"github.com/stretchr/testify/mock"
)

type IAPIKeyServices struct {
	mock.Mock
}

func (m *IAPIKeyServices) GenerateAPIKey(name, prefix, setup_id string, scope []string) (apikeys.APIKey, error) {

	args := m.Called(name, prefix, setup_id, scope)

	return args.Get(0).(apikeys.APIKey), args.Error(1)
}
