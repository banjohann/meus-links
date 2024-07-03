package api

import (
	"encoding/json"
	"net/http"
)

func DecodeBody(r *http.Request, v interface{}) error {
	err := json.NewDecoder(r.Body).Decode(v)
	return err
}
