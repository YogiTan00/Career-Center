package jobs

import (
	"CareerCenter/domain/entity"
	"CareerCenter/internal/delivery/request"
	response2 "CareerCenter/internal/delivery/response"
	"CareerCenter/utils"
	"context"
	"net/http"
)

func (u *JobsrHandler) GetListJob(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = context.TODO()
		req request.RequestFilter
	)
	_, errToken := utils.ValidateTokenFromHeader(r)
	if errToken != nil {
		http.Error(w, errToken.Error(), http.StatusUnauthorized)
		return
	}

	buildFilter := &entity.FilterDTO{
		Limit: req.Limit,
		Page:  req.Limit,
	}

	jobs, err := u.UCJobs.GetListJobs(ctx, buildFilter)
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
