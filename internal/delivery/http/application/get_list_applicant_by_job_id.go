package application

import (
	"CareerCenter/internal/delivery/response"
	"CareerCenter/logger"
	"CareerCenter/utils"
	"CareerCenter/utils/exceptions"
	"CareerCenter/utils/helper"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func (h *ApplicationHandler) GetApplicant(w http.ResponseWriter, r *http.Request) {
	var (
		ctx   = context.TODO()
		vars  = mux.Vars(r)
		jobId = vars["job_id"]
		log   = logger.NewLogger(fmt.Sprintf("/v1/admin/applicant/%s", jobId))
	)

	user, errToken := utils.ValidateTokenFromHeader(r)
	if errToken != nil {
		helper.ResponseErr(w, errToken, http.StatusUnauthorized)
		log.Error(errToken)
		return
	}

	if !user.Admin() {
		helper.ResponseErr(w, exceptions.Unauthorized, http.StatusUnauthorized)
		return
	}

	application, err := h.UCApplication.GetApplicantByJobId(ctx, jobId)
	if err != nil {
		helper.ResponseErr(w, err, http.StatusInternalServerError)
		log.Error(err)
		return
	}

	result := response.ResponseListApplicantByJobId(application, h.cfg)

	helper.ResponseInterface(w, "success get application by job id", result, http.StatusOK)
	log.InfoWithData("success get application by job id", nil)
	return
}
