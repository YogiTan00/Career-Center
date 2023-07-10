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

func (h *ProfileHandler) DeletedEducation(w http.ResponseWriter, r *http.Request) {
	var (
		ctx              = context.TODO()
		vars             = mux.Vars(r)
		workExperienceId = vars["education_id"]
		log              = logger.NewLogger(fmt.Sprintf("/v1/deleted-education/%s", workExperienceId))
	)

	_, errToken := utils.ValidateTokenFromHeader(r)
	if errToken != nil {
		helper.ResponseErr(w, errToken, http.StatusUnauthorized)
		log.General("", errToken)
		return
	}

	err := h.UCProfile.DeletedEducation(ctx, workExperienceId)
	if err != nil {
		helper.ResponseErr(w, err, http.StatusInternalServerError)
		log.General("", err)
		return
	}

	helper.Response(w, "success deleted education", http.StatusOK)
	log.General("success deleted education", nil)
	return
}
