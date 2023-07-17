package jobs

import (
	"CareerCenter/internal/delivery/request"
	"CareerCenter/logger"
	"CareerCenter/utils/helper"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func (h *JobsHandler) UpdateJob(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = context.TODO()
		vars    = mux.Vars(r)
		id      = vars["job_id"]
		req     *request.RequestJob
		decoder = json.NewDecoder(r.Body)
		log     = logger.NewLogger("/v1/admin/update/{job_id}")
	)
	errDecode := decoder.Decode(&req)
	if errDecode != nil {
		helper.ResponseErr(w, errDecode, http.StatusInternalServerError)
		log.Error(errDecode)
		return
	}

	buildUpdateJob := request.NewJobRequest(req)

	errUpdateUseCase := h.UCJobs.UpdateJob(ctx, id, buildUpdateJob)
	if errUpdateUseCase != nil {
		helper.ResponseErr(w, errUpdateUseCase, http.StatusInternalServerError)
		log.Error(errUpdateUseCase)
		return
	}

	helper.Response(w, "success update job", http.StatusOK)
	log.InfoWithData("Success update job", errUpdateUseCase)
	return
}
