package utils

import (
	"encoding/json"
	"net/http"
)

func DecodeBody(r *http.Request, e *json.Encoder, v interface{}) error {
	err := json.NewDecoder(r.Body).Decode(v)
	return err
}
