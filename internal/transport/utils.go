package transport

import (
	"encoding/json"
	"net/http"
)

func Respond(w http.ResponseWriter, data interface{}) error {
	w.Header().Add("Content-Type", "application/json")

	if data != nil {
		err := json.NewEncoder(w).Encode(data)

		if err != nil {
			return err
		}
	}

	return nil
}
