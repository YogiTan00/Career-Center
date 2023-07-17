package jobs

import (
	"CareerCenter/logger"
	"CareerCenter/utils/helper"
	"context"
	"github.com/gorilla/mux"
	"net/http"
)

func (h *JobsHandler) DeleteJob(w http.ResponseWriter, r *http.Request) {
	var (
		ctx  = context.TODO()
		vars = mux.Vars(r)
		id   = vars["job_id"]
		log  = logger.NewLogger("/v1/admin/delete/{job_id}")
	)

	errDeleteUseCase := h.UCJobs.DeleteJob(ctx, id)
	if errDeleteUseCase != nil {
		helper.ResponseErr(w, errDeleteUseCase, http.StatusInternalServerError)
		log.Error(errDeleteUseCase)
		return
	}

	helper.Response(w, "success delete job", http.StatusOK)
	log.InfoWithData("Success delete job", errDeleteUseCase)
	return
}
