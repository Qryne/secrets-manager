package v1

import (
	"encoding/json"
	"log"
	"net/http"

	apikeys "github.com/qryne/api/internal/api_keys"
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

func (ctrl *APIKeyController) CreateAPIKeyController(res http.ResponseWriter, req *http.Request) {
	var raw CreateAPIKeySchema
	err := json.NewDecoder(req.Body).Decode(&raw)

	if err != nil {
		http.Error(res, "Invalid request payload", http.StatusBadRequest)
		return
	}

	new_api_key, err := ctrl.APIKeysServices.GenerateAPIKey(raw.Name, raw.Prefix, raw.SetupID, raw.Scope)

	if err != nil {
		log.Fatal(err)
		http.Error(res, "Failed to create new API Key", http.StatusUnprocessableEntity)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(new_api_key)
}
