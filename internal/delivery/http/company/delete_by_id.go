package company

import (
	"CareerCenter/logger"
	"CareerCenter/utils/helper"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func (h *CompanyHandler) DeleteJob(w http.ResponseWriter, r *http.Request) {
	var (
		ctx       = context.TODO()
		vars      = mux.Vars(r)
		companyId = vars["company_id"]
		log       = logger.NewLogger(fmt.Sprintf("/v1/admin/delete-company/%s", companyId))
	)

	errDeleteUseCase := h.UCCompany.DeleteCompanyById(ctx, companyId)
	if errDeleteUseCase != nil {
		helper.ResponseErr(w, errDeleteUseCase, http.StatusInternalServerError)
		log.Error(errDeleteUseCase)
		return
	}

	helper.Response(w, "success delete company", http.StatusOK)
	log.InfoWithData("Success delete company", errDeleteUseCase)
	return
}
