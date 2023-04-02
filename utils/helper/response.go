package helper

import (
	"CareerCenter/internal/delivery/response"
	"net/http"
)

func ResponseErr(w http.ResponseWriter, err error, statusCode int) {
	result, errMap := response.MapResponse(1, err.Error())
	if errMap != nil {
		w.WriteHeader(statusCode)
		w.Write([]byte("Error mapping data"))
	}
	w.Write(result)
}
