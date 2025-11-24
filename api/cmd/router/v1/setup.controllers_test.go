package v1_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	v1 "github.com/qryne/api/cmd/router/v1"
	"github.com/qryne/api/internal/setups/mocks"
	"github.com/qryne/api/utility/responder"
	"github.com/stretchr/testify/assert"
)

type JsonResponse struct {
	Message string `json:"message"`
	Status  string `json:"success"`
	Data    any    `json:"data"`
}

func TestInitSetupController(t *testing.T) {
	setupService := new(mocks.ISetupServices)

	setupService.On("InitSetup").Return(nil)

	body := strings.NewReader(`{}`)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/setup", body)
	w := httptest.NewRecorder()

	controller := v1.SetupController{SetupServices: setupService}

	r := chi.NewRouter()
	r.HandleFunc("/api/v1/setup", controller.InitSetupController)
	r.ServeHTTP(w, req)

	var result responder.JsonResponse[any]
	_ = json.NewDecoder(w.Body).Decode(&result)

	assert.Equal(t, "Setup initiated successfully", result.Message)
	assert.Equal(t, responder.StatusSuccess, result.Status)
	assert.Nil(t, result.Data)
}
