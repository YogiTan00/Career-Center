package jobs

import (
	"CareerCenter/internal/delivery/request"
	"CareerCenter/internal/delivery/response"
	"CareerCenter/utils"
	"CareerCenter/utils/helper"
	"context"
	"net/http"
)

func (u *JobsHandler) GetListJob(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = context.TODO()
		req request.RequestFilter
	)

	_, errToken := utils.ValidateTokenFromHeader(r)
	if errToken != nil {
		helper.ResponseErr(w, errToken, http.StatusUnauthorized)
		return
	}

	filter, errFilter := request.FilterGeneral(r, &req)
	if errFilter != nil {
		helper.ResponseErr(w, errFilter, http.StatusBadRequest)
		return
	}

	jobs, err := u.UCJobs.GetListJobs(ctx, filter)
	if err != nil {
		helper.ResponseErr(w, err, http.StatusInternalServerError)
		return
	}

	JobsResponse := response.GetListJobResponse(jobs)
	helper.ResponseInterface(w, "success Get list job", JobsResponse, http.StatusInternalServerError)
	return
}
