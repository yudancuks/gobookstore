package utils

import (
	"books-list/models"
	"encoding/json"
	"net/http"
)

// SendErr : Send the error message to controller if there is error
func SendErr(w http.ResponseWriter, status int, err models.Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(err)
}

// SendSuccess : Send the data to controller
func SendSuccess(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}
