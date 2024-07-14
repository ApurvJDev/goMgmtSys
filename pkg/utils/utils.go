package utils

// We need data that is unmarshaled
// we do not need JSON data
// so we need to unmarshal it to use it in our controller
// this is the purpose of this file

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
