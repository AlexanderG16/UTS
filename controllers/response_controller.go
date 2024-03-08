package controllers

import (
	m "UTS/models"
	"encoding/json"
	"net/http"
)

func SendSuccessResponse(w http.ResponseWriter, status int, message string) {
	var response m.GlobalResponse
	response.Status = status
	response.Message = message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func SendErrorResponse(w http.ResponseWriter, status int, message string) {
	var response m.GlobalResponse
	response.Status = status
	response.Message = message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
