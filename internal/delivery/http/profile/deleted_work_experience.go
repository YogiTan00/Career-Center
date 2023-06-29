package profile

import (
	"CareerCenter/logger"
	"CareerCenter/utils"
	"CareerCenter/utils/helper"
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *ProfileHandler) DeletedWorkExperience(w http.ResponseWriter, r *http.Request) {
	var (
		ctx              = context.TODO()
		vars             = mux.Vars(r)
		workExperienceId = vars["work_experience_id"]
		log              = logger.NewLogger(fmt.Sprintf("/v1/deleted-work-experience/%s", workExperienceId))
	)

	_, errToken := utils.ValidateTokenFromHeader(r)
	if errToken != nil {
		helper.ResponseErr(w, errToken, http.StatusUnauthorized)
		log.General("", errToken)
		return
	}

	err := h.UCProfile.DeletedWorkExperience(ctx, workExperienceId)
	if err != nil {
		helper.ResponseErr(w, err, http.StatusInternalServerError)
		log.General("", err)
		return
	}

	helper.Response(w, "success deleted work experience", http.StatusInternalServerError)
	log.General("success deleted work experience", nil)
	return
}
