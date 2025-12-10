package v1

import (
	"encoding/json"
	"log/slog"
	"net/http"

	apikeys "github.com/qryne/api/internal/api_keys"
	"github.com/qryne/api/utility/responder"
)

type APIKeyController struct {
	APIKeysServices apikeys.IAPIKeyServices
}

type CreateAPIKeySchema struct {
	Name    string   `json:"name"`
	Prefix  string   `json:"prefix"`
	SetupID string   `json:"setup_id"`
	Scope   []string `json:"scope"`
}

func (ctrl *APIKeyController) CreateAPIKeyController(W http.ResponseWriter, R *http.Request) {
	var raw CreateAPIKeySchema
	err := json.NewDecoder(R.Body).Decode(&raw)

	if err != nil {
		resp := responder.NewFailed[any]("Invalid request payload", nil)
		responder.WriteJSON(W, http.StatusBadRequest, resp)
		return
	}

	new_api_key, err := ctrl.APIKeysServices.GenerateAPIKey(raw.Name, raw.Prefix, raw.SetupID, raw.Scope)
	if err != nil {
		slog.Error("Failed to create API key:", err)
		resp := responder.NewFailed[any]("Failed to generate API key", nil)
		responder.WriteJSON(W, http.StatusUnprocessableEntity, resp)
		return
	}

	resp := responder.NewSuccess("API key generated successfully", &new_api_key)
	responder.WriteJSON(W, http.StatusCreated, resp)
}
