package api

import (
	"encoding/json"
	"net/http"
)

type HttpError struct {
	Message string `json:"message"`
	Detail  string `json:"detail"`
	Code    int    `json:"code"`
	Path    string `json:"path"`
}

func newHttpError(message, detail, path string, code int) *HttpError {
	return &HttpError{
		Message: message,
		Detail:  detail,
		Path:    path,
		Code:    code,
	}
}

func (error *HttpError) Write(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	w.WriteHeader(error.Code)
	encoder.Encode(error)
}

func ErrorBadRequest(message, detail string, w http.ResponseWriter, r *http.Request) {
	newHttpError(message, detail, r.URL.Path, 400).Write(w, r)
}

func ErrorUnauthorized(message, detail string, w http.ResponseWriter, r *http.Request) {
	newHttpError(message, detail, r.URL.Path, 401).Write(w, r)
}

func ErrorNotFound(message, detail string, w http.ResponseWriter, r *http.Request) {
	newHttpError(message, detail, r.URL.Path, 404).Write(w, r)
}

func ErrorInternal(message, detail string, w http.ResponseWriter, r *http.Request) {
	newHttpError(message, detail, r.URL.Path, 500).Write(w, r)
}
