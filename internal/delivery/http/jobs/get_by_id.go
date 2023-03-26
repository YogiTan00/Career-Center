package jobs

import (
	response2 "CareerCenter/internal/delivery/response"
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
		http.Error(w, errToken.Error(), http.StatusUnauthorized)
		return
	}

	jobs, err := h.UCJobs.GetJobById(ctx, jobId)
	if err != nil {
		response, errMap := response2.MapResponse(1, "cant get detail job")
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}
		w.Write(response)
	} else {
		response := response2.GetDetailJobResponse(jobs)
		result, errMap := response2.MapResponseInterface(0, "success Get detail job", response)
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}
		w.Write(result)
	}
}
