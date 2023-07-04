package helper

import (
	"CareerCenter/internal/delivery/response"
	"net/http"
)

func Response(w http.ResponseWriter, msg string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	result, errMap := response.MapResponse(0, msg)
	if errMap != nil {
		w.WriteHeader(statusCode)
		w.Write([]byte("Error mapping data"))
	}
	w.Write(result)
}

func ResponseInterface(w http.ResponseWriter, msg string, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	result, errMap := response.MapResponseInterface(0, msg, data)
	if errMap != nil {
		w.WriteHeader(statusCode)
		w.Write([]byte("Error mapping data"))
	}
	w.Write(result)
}

func ResponseErr(w http.ResponseWriter, err error, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	result, errMap := response.MapResponse(1, err.Error())
	if errMap != nil {
		w.WriteHeader(statusCode)
		w.Write([]byte("Error mapping data"))
	}
	http.Error(w, "", statusCode)
	w.Write(result)
}
