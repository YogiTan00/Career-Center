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
		log = logger.NewLogger(r.RequestURI)
	)

	_, errToken := utils.ValidateTokenFromHeader(r)
	if errToken != nil {
		helper.ResponseErr(w, errToken, http.StatusUnauthorized)
		log.General("", errToken)
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
		log.General("", err)
		return
	}

	JobsResponse := response.GetListJobResponse(jobs)
	helper.ResponseInterface(w, "success Get list job", JobsResponse, http.StatusInternalServerError)
	log.General("success send application", JobsResponse)
	return
}
