package utils

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	StatusCode    int    `json:"status"`
	ErrorMessage  string `json:"message"`
}

func CreateErrorResponse(err error, statusCode int, w http.ResponseWriter) {
	resp := &ErrorResponse{statusCode, err.Error()}
	respJSON, _ := json.Marshal(resp)
	w.WriteHeader(statusCode)
	w.Write(respJSON)
}

type SuccessResponse struct {
	StatusCode int         `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func CreateSuccessResponse(message string, data interface{}, statusCode int, w http.ResponseWriter) {
	resp := &SuccessResponse{statusCode, message, data}
	respJSON, _ := json.Marshal(resp)
	w.WriteHeader(statusCode)
	w.Write((respJSON))
}

func CreateLoginJwtToken(){
	
}