package profile

import (
	"CareerCenter/logger"
	"CareerCenter/utils"
	"CareerCenter/utils/helper"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
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
		log.Error(errToken)
		return
	}

	err := h.UCProfile.DeletedWorkExperience(ctx, workExperienceId)
	if err != nil {
		helper.ResponseErr(w, err, http.StatusInternalServerError)
		log.Error(err)
		return
	}

	helper.Response(w, "success deleted work experience", http.StatusOK)
	log.InfoWithData("success deleted work experience", nil)
	return
}
