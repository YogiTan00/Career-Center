package jobs

import (
	"CareerCenter/internal/delivery/request"
	"CareerCenter/internal/delivery/response"
	"CareerCenter/logger"
	"CareerCenter/utils"
	"CareerCenter/utils/helper"
	"context"
	"net/http"
)

func (u *JobsHandler) GetListJob(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = context.TODO()
		req request.RequestFilter
		log = logger.NewLogger("/v1/list-jobs")
	)

	_, errToken := utils.ValidateTokenFromHeader(r)
	if errToken != nil {
		helper.ResponseErr(w, errToken, http.StatusUnauthorized)
		log.Error(errToken)
		return
	}

	filter, errFilter := request.FilterGeneral(r, &req)
	if errFilter != nil {
		helper.ResponseErr(w, errFilter, http.StatusBadRequest)
		return
	}

	jobs, count, err := u.UCJobs.GetListJobs(ctx, filter)
	if err != nil {
		helper.ResponseErr(w, err, http.StatusInternalServerError)
		log.Error(err)
		return
	}

	JobsResponse := response.GetListJobResponse(jobs)
	helper.ResponseInterfaceWithCount(w, "success get list job", JobsResponse, count, http.StatusOK)
	log.InfoWithData("success send application", JobsResponse)
	return
}
