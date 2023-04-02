package jobs

import (
	"CareerCenter/internal/delivery/request"
	"CareerCenter/internal/delivery/response"
	"CareerCenter/utils"
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
		result, errMap := response.MapResponse(1, errToken.Error())
		if errMap != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Error mapping data"))
		}
		w.Write(result)
		return
	}

	filter, errFilter := request.FilterGeneral(r, &req)
	if errFilter != nil {
		result, errMap := response.MapResponse(1, errFilter.Error())
		if errMap != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error mapping data"))
		}
		w.Write(result)
		return
	}

	jobs, err := u.UCJobs.GetListJobs(ctx, filter)
	if err != nil {
		result, errMap := response.MapResponse(1, "cant get list job")
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}
		w.Write(result)
		return
	}

	JobsResponse := response.GetListJobResponse(jobs)
	result, errMap := response.MapResponseInterface(0, "success Get list job", JobsResponse)
	if errMap != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error mapping data"))
	}
	w.Write(result)
	return
}
