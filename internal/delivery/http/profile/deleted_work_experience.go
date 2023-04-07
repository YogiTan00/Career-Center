package profile

import (
	"CareerCenter/internal/delivery/response"
	"CareerCenter/utils"
	"CareerCenter/utils/helper"
	"context"
	"github.com/gorilla/mux"
	"net/http"
)

func (h *ProfileHandler) DeletedWorkExperience(w http.ResponseWriter, r *http.Request) {
	var (
		ctx  = context.TODO()
		vars = mux.Vars(r)
	)

	workExperienceId := vars["work_experience_id"]

	_, errToken := utils.ValidateTokenFromHeader(r)
	if errToken != nil {
		helper.ResponseErr(w, errToken, http.StatusUnauthorized)
		return
	}

	err := h.UCProfile.DeletedWorkExperience(ctx, workExperienceId)
	if err != nil {
		helper.ResponseErr(w, err, http.StatusInternalServerError)
		return
	}

	result, errMap := response.MapResponse(0, "success deleted work experience")
	if errMap != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error mapping data"))
	}
	w.Write(result)
	return
}
