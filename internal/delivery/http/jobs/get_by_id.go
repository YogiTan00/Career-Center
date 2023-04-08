package jobs

import (
	"CareerCenter/internal/delivery/response"
	"CareerCenter/utils"
	"CareerCenter/utils/helper"
	"context"
	"github.com/gorilla/mux"
	"net/http"
)

func (h *JobsHandler) GetJobById(w http.ResponseWriter, r *http.Request) {
	var (
		ctx  = context.TODO()
		vars = mux.Vars(r)
	)

	jobId := vars["job_id"]

	email, errToken := utils.ValidateTokenFromHeader(r)
	if errToken != nil {
		helper.ResponseErr(w, errToken, http.StatusUnauthorized)
		return
	}

	jobs, err := h.UCJobs.GetJobById(ctx, email, jobId)
	if err != nil {
		helper.ResponseErr(w, err, http.StatusInternalServerError)
		return
	}

	jobsResponse := response.GetDetailJobResponse(jobs)
	helper.ResponseInterface(w, "success Get detail job", jobsResponse, http.StatusInternalServerError)
	return
}
