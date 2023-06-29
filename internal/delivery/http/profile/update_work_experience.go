package profile

import (
	"CareerCenter/internal/delivery/request"
	"CareerCenter/logger"
	"CareerCenter/utils"
	"CareerCenter/utils/helper"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *ProfileHandler) UpdateWorkExperience(w http.ResponseWriter, r *http.Request) {
	var (
		ctx              = context.TODO()
		req              *request.RequestWorkExperience
		decoder          = json.NewDecoder(r.Body)
		vars             = mux.Vars(r)
		workExperienceId = vars["work_experience_id"]
		log              = logger.NewLogger(fmt.Sprintf("/v1/update-work-experience/%s", workExperienceId))
	)

	errDecode := decoder.Decode(&req)
	if errDecode != nil {
		helper.ResponseErr(w, errDecode, http.StatusUnauthorized)
		log.General("", errDecode)
		return
	}

	_, errToken := utils.ValidateTokenFromHeader(r)
	if errToken != nil {
		helper.ResponseErr(w, errToken, http.StatusUnauthorized)
		log.General("", errToken)
		return
	}

	workExperience, err := request.NewUpdateWorkExperience(req)
	if err != nil {
		helper.ResponseErr(w, err, http.StatusBadRequest)
		log.General("", err)
		return
	}

	err = h.UCProfile.UpdateWorkExperience(ctx, workExperienceId, workExperience)
	if err != nil {
		helper.ResponseErr(w, err, http.StatusInternalServerError)
		log.General("", err)
		return
	}

	helper.Response(w, "success update work experience", http.StatusInternalServerError)
	log.General("success update work experience", nil)
	return
}
