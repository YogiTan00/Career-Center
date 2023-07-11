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
		_, err := w.Write([]byte("Error mapping data"))
		if err != nil {
			return
		}
	}
	_, err := w.Write(result)
	if err != nil {
		return
	}
}

func ResponseInterface(w http.ResponseWriter, msg string, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	result, errMap := response.MapResponseInterface(0, msg, data)
	if errMap != nil {
		w.WriteHeader(statusCode)
		_, err := w.Write([]byte("Error mapping data"))
		if err != nil {
			return
		}
	}
	_, err := w.Write(result)
	if err != nil {
		return
	}
}

func ResponseInterfaceWithCount(w http.ResponseWriter, msg string, data interface{}, count int, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	result, errMap := response.MapResponseInterfaceWithCount(0, msg, data, count)
	if errMap != nil {
		w.WriteHeader(statusCode)
		_, err := w.Write([]byte("Error mapping data"))
		if err != nil {
			return
		}
	}
	_, err := w.Write(result)
	if err != nil {
		return
	}
}

func ResponseErr(w http.ResponseWriter, err error, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	result, errMap := response.MapResponse(1, err.Error())
	if errMap != nil {
		w.WriteHeader(statusCode)
		_, err := w.Write([]byte("Error mapping data"))
		if err != nil {
			return
		}
	}
	http.Error(w, "", statusCode)
	_, err = w.Write(result)
	if err != nil {
		return
	}
}
