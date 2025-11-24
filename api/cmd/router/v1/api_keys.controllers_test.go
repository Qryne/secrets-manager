package v1_test

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	v1 "github.com/qryne/api/cmd/router/v1"
	apikeys "github.com/qryne/api/internal/api_keys"
	"github.com/qryne/api/internal/api_keys/mocks"
	"github.com/stretchr/testify/assert"
)

const API_KEY_NAME = "Global API Key"
const API_KEY_SLUG = "global-api-key"
const SETUP_ID = "f33225a7-9d93-4c8b-b545-9403a298e08e"

func TestControllerAPIKeyController(t *testing.T) {
	apiKeyServices := new(mocks.IAPIKeyServices)

	apiKeyServices.On("GenerateAPIKey", API_KEY_NAME, SETUP_ID, []string{"super_admin", "owner"})

	controller := v1.APIKeyController{APIKeysServices: apiKeyServices}

	body := strings.NewReader(`{"name": %s, "setup_id": %s, "scope": ["super_admin", "owner"]}`)
	req := httptest.NewRequest("POST", "/api/v1/api_keys", body)
	w := httptest.NewRecorder()

	r := chi.NewRouter()
	r.HandleFunc("/api/v1/api_keys", controller.CreateAPIKeyController)

	r.ServeHTTP(w, req)

	expectedResult := apikeys.APIKey{}
	expectedResult.Name = API_KEY_NAME
	expectedResult.Algorithm = "AES256"
	expectedResult.Slug = API_KEY_SLUG

	actualResult := apikeys.APIKey{}
	json.NewDecoder(w.Body).Decode(&actualResult)

	assert.Equal(t, expectedResult, actualResult)
}
