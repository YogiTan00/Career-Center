package jobs

import (
	"CareerCenter/internal/delivery/request"
	response2 "CareerCenter/internal/delivery/response"
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
		http.Error(w, errToken.Error(), http.StatusUnauthorized)
		return
	}

	filter, errFilter := request.FilterGeneral(r, &req)
	if errFilter != nil {
		response, errMap := response2.MapResponse(1, errFilter.Error())
		if errMap != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error mapping data"))
		}
		w.Write(response)
		return
	}

	jobs, err := u.UCJobs.GetListJobs(ctx, filter)
	if err != nil {
		response, errMap := response2.MapResponse(1, "wrong email or password")
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}
		w.Write(response)
	} else {
		response := response2.GetListJob(jobs)
		result, errMap := response2.MapResponseInterface(0, "success login", response)
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}
		w.Write(result)
	}
}
