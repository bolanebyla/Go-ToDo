package transport

import (
	"encoding/json"
	"net/http"
)

func Response() map[string]interface{} {
	return map[string]interface{}{}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) error {
	w.Header().Add("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(data)

	if err != nil {
		return err
	}
	return nil
}
