package v1_test

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	v1 "github.com/qryne/api/cmd/router/v1"
	apikeys "github.com/qryne/api/internal/api_keys"
	"github.com/qryne/api/internal/api_keys/mocks"
	"github.com/qryne/api/utility/responder"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const API_KEY_NAME = "Global API Key"
const API_KEY_SLUG = "global-api-key"
const PREFIX = "SAT"
const SETUP_ID = "f33225a7-9d93-4c8b-b545-9403a298e08e"

func TestControllerAPIKeyController(t *testing.T) {
	apiKeyServices := new(mocks.IAPIKeyServices)

	expectedResult := apikeys.APIKey{
		Name:   API_KEY_NAME,
		Prefix: PREFIX,
		Scope:  []string{"super_admin", "owner"},
	}
	apiKeyServices.On("GenerateAPIKey", API_KEY_NAME, PREFIX, SETUP_ID, mock.AnythingOfType("[]string")).Return(expectedResult, nil)

	controller := v1.APIKeyController{APIKeysServices: apiKeyServices}

	body := strings.NewReader(fmt.Sprintf(`{"name": "%s", "prefix": "%s", "setup_id": "%s", "scope": ["super_admin", "owner"]}`, API_KEY_NAME, PREFIX, SETUP_ID))
	req := httptest.NewRequest("POST", "/api/v1/api_keys", body)
	w := httptest.NewRecorder()

	r := chi.NewRouter()
	r.HandleFunc("/api/v1/api_keys", controller.CreateAPIKeyController)

	r.ServeHTTP(w, req)

	var result responder.JsonResponse[apikeys.APIKey]
	_ = json.NewDecoder(w.Body).Decode(&result)

	assert.Equal(t, "API key generated successfully", result.Message)
	assert.Equal(t, responder.StatusSuccess, result.Status)
	assert.Equal(t, expectedResult.Name, result.Data.Name)
	assert.Equal(t, expectedResult.Scope, result.Data.Scope)
}
