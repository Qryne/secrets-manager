package responder

import (
	"encoding/json"
	"net/http"
)

// Status values
const (
	StatusSuccess = "SUCCESS"
	StatusFailed  = "FAILED"
)

type JsonResponse[T any] struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Data    *T     `json:"data"`
}

func NewSuccess[T any](message string, data *T) JsonResponse[T] {
	return JsonResponse[T]{Message: message, Status: StatusSuccess, Data: data}
}

func NewFailed[T any](message string, data *T) JsonResponse[T] {
	return JsonResponse[T]{Message: message, Status: StatusFailed, Data: data}
}

func WriteJSON[T any](w http.ResponseWriter, code int, resp JsonResponse[T]) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(resp)
}
