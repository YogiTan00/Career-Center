package profile

import (
	"CareerCenter/internal/delivery/request"
	"CareerCenter/internal/delivery/response"
	"CareerCenter/utils"
	"CareerCenter/utils/helper"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func (h *ProfileHandler) UpdateWorkExperience(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = context.TODO()
		req     *request.RequestWorkExperience
		decoder = json.NewDecoder(r.Body)
		vars    = mux.Vars(r)
	)

	workExperienceId := vars["work_experience_id"]

	errDecode := decoder.Decode(&req)
	if errDecode != nil {
		helper.ResponseErr(w, errDecode, http.StatusUnauthorized)
		return
	}

	_, errToken := utils.ValidateTokenFromHeader(r)
	if errToken != nil {
		helper.ResponseErr(w, errToken, http.StatusUnauthorized)
		return
	}

	workExperience, err := request.NewUpdateWorkExperience(req)
	if err != nil {
		helper.ResponseErr(w, err, http.StatusBadRequest)
		return
	}

	err = h.UCProfile.UpdateWorkExperience(ctx, workExperienceId, workExperience)
	if err != nil {
		helper.ResponseErr(w, err, http.StatusInternalServerError)
		return
	}

	result, errMap := response.MapResponse(0, "success update profile")
	if errMap != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error mapping data"))
	}
	w.Write(result)
	return
}
