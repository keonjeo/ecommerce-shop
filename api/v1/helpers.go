package apiv1

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func respondJSON(w http.ResponseWriter, code int, obj interface{}) error {
	b, err := json.Marshal(obj)
	if err != nil {
		return fmt.Errorf("could not encode json response: %v - %v", obj, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(b)
	return err
}

func respondError() {}

func respondOK() {}
