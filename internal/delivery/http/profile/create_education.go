package profile

import (
	"CareerCenter/internal/delivery/request"
	"CareerCenter/internal/delivery/response"
	"CareerCenter/utils"
	"context"
	"encoding/json"
	"net/http"
)

func (h *ProfileHandler) CreateEducation(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = context.TODO()
		req     *request.RequestEducation
		decoder = json.NewDecoder(r.Body)
	)
	errDecode := decoder.Decode(&req)
	if errDecode != nil {
		result, errMap := response.MapResponse(1, errDecode.Error())
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}
		w.Write(result)
		return
	}

	email, errToken := utils.ValidateTokenFromHeader(r)
	if errToken != nil {
		result, errMap := response.MapResponse(1, errToken.Error())
		if errMap != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Error mapping data"))
		}
		w.Write(result)
		return
	}

	education, err := request.NewUpdateEducation(req)
	if err != nil {
		result, errMap := response.MapResponse(1, errDecode.Error())
		if errMap != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error mapping data"))
		}
		w.Write(result)
		return
	}

	err = h.UCProfile.CreateEducation(ctx, email, education)
	if err != nil {
		result, errMap := response.MapResponse(1, err.Error())
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}
		w.Write(result)
		return
	}

	result, errMap := response.MapResponse(0, "success add work experience")
	if errMap != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error mapping data"))
	}
	w.Write(result)
	return
}
