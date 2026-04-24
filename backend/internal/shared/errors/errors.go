package errors

import (
	"encoding/json"
	"net/http"
)

type APIError struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func WriteJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func WriteError(w http.ResponseWriter, status int, code, message string) {
	e := APIError{}
	e.Error.Code = code
	e.Error.Message = message
	WriteJSON(w, status, e)
}
