package apikeys_test

import (
	"testing"

	apikeys "github.com/qryne/api/internal/api_keys"
	"github.com/qryne/api/internal/api_keys/mocks"
	db_gen "github.com/qryne/api/internal/db/sqlc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGenerateAPIKey(t *testing.T) {
	t.Setenv("SETUP_API_SECRET", "8WPNbbV1lOpGKqzbsGir7A==")

	apiKeyRepo := new(mocks.IAPIKeyRepository)

	scope := []string{"super_admin", "owner"}

	input := apikeys.APIKey{
		Name:   "Global API Key",
		Slug:   "global-api-key",
		Prefix: "SKEY",
		Scope:  scope,
	}

	expected := db_gen.ApiKey{
		Name:      input.Name,
		Prefix:    input.Prefix,
		Scope:     scope,
		Algorithm: "AES256",
	}

	// Repository expectation
	apiKeyRepo.On(
		"CreateAPIKey",
		input.Name,
		input.Slug,
		input.Prefix,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		expected.Algorithm,
		"f33225a7-9d93-4c8b-b545-9403a298e08e",
		scope,
	).Return(expected, nil)

	service := apikeys.APIKeyServices{APIKeyRepo: apiKeyRepo}

	result, err := service.GenerateAPIKey(
		input.Name,
		input.Prefix,
		"f33225a7-9d93-4c8b-b545-9403a298e08e",
		scope,
	)

	assert.NoError(t, err)
	assert.Equal(t, expected.Name, result.Name)
	assert.Equal(t, expected.Prefix, result.Prefix)
	assert.Equal(t, expected.Scope, result.Scope)
	apiKeyRepo.AssertExpectations(t)
}
