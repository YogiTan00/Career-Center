package profile

import (
	"CareerCenter/internal/delivery/request"
	"CareerCenter/internal/delivery/response"
	"CareerCenter/utils"
	"CareerCenter/utils/helper"
	"context"
	"encoding/json"
	"net/http"
)

func (h *ProfileHandler) CreateWorkExperience(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = context.TODO()
		req     *request.RequestWorkExperience
		decoder = json.NewDecoder(r.Body)
	)
	errDecode := decoder.Decode(&req)
	if errDecode != nil {
		helper.ResponseErr(w, errDecode, http.StatusUnauthorized)
		return
	}

	email, errToken := utils.ValidateTokenFromHeader(r)
	if errToken != nil {
		helper.ResponseErr(w, errToken, http.StatusUnauthorized)
		return
	}

	workExperience, err := request.NewUpdateWorkExperience(req)
	if err != nil {
		helper.ResponseErr(w, err, http.StatusBadRequest)
		return
	}

	err = h.UCProfile.CreateWorkExperiencet(ctx, email, workExperience)
	if err != nil {
		helper.ResponseErr(w, err, http.StatusInternalServerError)
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
