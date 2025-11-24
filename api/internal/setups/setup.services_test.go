package setups_test

import (
	"testing"

	db_gen "github.com/qryne/api/internal/db/sqlc"
	"github.com/qryne/api/internal/setups"
	"github.com/qryne/api/internal/setups/mocks"
	"github.com/stretchr/testify/assert"
)

func TestInitSetup(t *testing.T) {
	setupRepo := new(mocks.ISetupRepository)

	expectedResult := db_gen.Setup{
		IsSetupComplete: false,
	}

	setupRepo.On("CreateEntry").Return(expectedResult, nil)

	service := setups.SetupServices{SetupRepo: setupRepo}

	err := service.InitSetup()

	assert.NoError(t, err)
	setupRepo.AssertExpectations(t)
}
