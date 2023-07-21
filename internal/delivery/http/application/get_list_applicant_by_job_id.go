package application

import (
	"CareerCenter/logger"
	"CareerCenter/utils"
	"CareerCenter/utils/exceptions"
	"CareerCenter/utils/helper"
	"context"
	"github.com/gorilla/mux"
	"net/http"
)

func (h *ApplicationHandler) GetApplicant(w http.ResponseWriter, r *http.Request) {
	var (
		ctx   = context.TODO()
		vars  = mux.Vars(r)
		jobId = vars["job_id"]
		log   = logger.NewLogger("/v1/admin/applicant/{job_id}")
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

	helper.ResponseInterface(w, "success get application by job id", application, http.StatusOK)
	log.InfoWithData("success get application by job id", nil)
	return
}
