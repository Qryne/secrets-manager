package v1

import (
	"log/slog"
	"net/http"

	"github.com/qryne/api/internal/setups"
	"github.com/qryne/api/utility/responder"
)

type SetupController struct {
	SetupServices setups.ISetupServices
}

func (controller *SetupController) InitSetupController(W http.ResponseWriter, R *http.Request) {

	err := controller.SetupServices.InitSetup()
	if err != nil {
		slog.Error("Setup Error:", err)
		resp := responder.NewFailed[any]("Failed to initiate setup", nil)
		responder.WriteJSON(W, http.StatusBadRequest, resp)
		return
	}

	resp := responder.NewSuccess[any]("Setup initiated successfully", nil)
	responder.WriteJSON(W, http.StatusCreated, resp)
}
