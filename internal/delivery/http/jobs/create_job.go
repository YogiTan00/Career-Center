package jobs

import (
	"CareerCenter/internal/delivery/request"
	"CareerCenter/logger"
	"CareerCenter/utils/helper"
	"context"
	"encoding/json"
	"net/http"
)

func (h *JobsHandler) CreateJob(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = context.TODO()
		req     *request.RequestJob
		decoder = json.NewDecoder(r.Body)
		log     = logger.NewLogger("/v1/admin/job")
	)
	errDecode := decoder.Decode(&req)
	if errDecode != nil {
		helper.ResponseErr(w, errDecode, http.StatusInternalServerError)
		log.General("", errDecode)
		return
	}

	buildCreateJob := request.NewJobRequest(req)

	errRegisterUseCase := h.UCJobs.CreateJob(ctx, buildCreateJob)
	if errRegisterUseCase != nil {
		helper.ResponseErr(w, errRegisterUseCase, http.StatusInternalServerError)
		log.General("", errRegisterUseCase)
		return
	}

	helper.Response(w, "success create job", http.StatusOK)
	log.General("Success create job", errRegisterUseCase)
	return
}
