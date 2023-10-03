package jobs

import (
	"CareerCenter/internal/delivery/response"
	"CareerCenter/logger"
	"CareerCenter/utils"
	"CareerCenter/utils/helper"
	"context"
	"net/http"
)

func (u *JobsHandler) GetListJobByEmail(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = context.TODO()
		log = logger.NewLogger("/v1/my-application")
	)

	user, errToken := utils.ValidateTokenFromHeader(r)
	if errToken != nil {
		helper.ResponseErr(w, errToken, http.StatusUnauthorized)
		log.Error(errToken)
		return
	}

	jobs, count, err := u.UCJobs.GetListByEmail(ctx, user)
	if err != nil {
		helper.ResponseErr(w, err, http.StatusInternalServerError)
		log.Error(err)
		return
	}

	JobsResponse := response.GetListJobResponse(jobs, u.cfg)
	helper.ResponseInterfaceWithCount(w, "success get list job", JobsResponse, count, http.StatusOK)
	log.InfoWithData("success send application", JobsResponse)
	return
}
