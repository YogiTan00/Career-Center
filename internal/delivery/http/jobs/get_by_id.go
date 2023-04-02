package jobs

import (
	"CareerCenter/internal/delivery/response"
	"CareerCenter/utils"
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

	jobs, err := h.UCJobs.GetJobById(ctx, jobId)
	if err != nil {
		result, errMap := response.MapResponse(1, "cant get detail job")
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}
		w.Write(result)
		return
	}

	jobsResponse := response.GetDetailJobResponse(jobs)
	result, errMap := response.MapResponseInterface(0, "success Get detail job", jobsResponse)
	if errMap != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error mapping data"))
	}
	w.Write(result)
	return
}
