package profile

import (
	"CareerCenter/internal/delivery/request"
	"CareerCenter/logger"
	"CareerCenter/utils"
	"CareerCenter/utils/helper"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func (h *ProfileHandler) UpdateEducation(w http.ResponseWriter, r *http.Request) {
	var (
		ctx              = context.TODO()
		req              *request.RequestEducation
		decoder          = json.NewDecoder(r.Body)
		vars             = mux.Vars(r)
		workExperienceId = vars["education_id"]
		log              = logger.NewLogger(fmt.Sprintf("/v1/profile/update-education/%s", workExperienceId))
	)

	errDecode := decoder.Decode(&req)
	if errDecode != nil {
		helper.ResponseErr(w, errDecode, http.StatusInternalServerError)
		log.General("", errDecode)
		return
	}

	_, errToken := utils.ValidateTokenFromHeader(r)
	if errToken != nil {
		helper.ResponseErr(w, errToken, http.StatusUnauthorized)
		log.General("", errToken)
		return
	}

	workExperience, err := request.NewUpdateEducation(req)
	if err != nil {
		helper.ResponseErr(w, err, http.StatusBadRequest)
		log.General("", err)
		return
	}

	err = h.UCProfile.UpdateEducation(ctx, workExperienceId, workExperience)
	if err != nil {
		helper.ResponseErr(w, err, http.StatusInternalServerError)
		log.General("", err)
		return
	}

	helper.Response(w, "success update education", http.StatusOK)
	log.General("success update education", nil)
	return
}
