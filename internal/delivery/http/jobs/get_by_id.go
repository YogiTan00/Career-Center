package jobs

import (
	"CareerCenter/internal/delivery/response"
	"CareerCenter/logger"
	"CareerCenter/utils"
	"CareerCenter/utils/helper"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func (h *JobsHandler) GetJobById(w http.ResponseWriter, r *http.Request) {
	var (
		ctx   = context.TODO()
		vars  = mux.Vars(r)
		jobId = vars["job_id"]
		log   = logger.NewLogger(fmt.Sprintf("/v1/job-detail/%s", jobId))
	)

	user, errToken := utils.ValidateTokenFromHeader(r)
	if errToken != nil {
		log.Error(errToken)
		helper.ResponseErr(w, errToken, http.StatusUnauthorized)
		return
	}

	jobs, err := h.UCJobs.GetJobById(ctx, user, jobId)
	if err != nil {
		helper.ResponseErr(w, err, http.StatusInternalServerError)
		log.Error(err)
		return
	}

	jobsResponse := response.GetDetailJobResponse(jobs)
	helper.ResponseInterface(w, "success get detail job", jobsResponse, http.StatusOK)
	log.InfoWithData("success Get detail job", jobsResponse)
	return
}
