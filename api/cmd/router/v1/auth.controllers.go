package v1

import (
	"encoding/json"
	"net/http"

	"github.com/qryne/api/internal/auth"
)

type AuthController struct {
	AuthServices auth.IAuthServices
}

type UserSignupSchema struct {
	Email string `json:"email"`
}

func (ctrl *AuthController) UserSignup(res http.ResponseWriter, req *http.Request) {
	var raw UserSignupSchema
	err := json.NewDecoder(req.Body).Decode(&raw)

	if err != nil {
		http.Error(res, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = ctrl.AuthServices.InitUserSignup(raw.Email)

	if err != nil {
		http.Error(res, "Failed to signup", http.StatusUnprocessableEntity)
	}

	res.WriteHeader(http.StatusCreated)
}
