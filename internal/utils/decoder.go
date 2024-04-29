package utils

import (
	"encoding/json"
	"net/http"

	"github.com/JohannBandelow/meus-links-go/internal/api"
)

func DecodeBody(r *http.Request, e *json.Encoder, v interface{}) {
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		e.Encode(api.ErrorBadRequest("Invalid request payload"))
		return
	}
}
