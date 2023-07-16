package company

import (
	"CareerCenter/internal/delivery/response"
	"CareerCenter/logger"
	"CareerCenter/utils"
	"CareerCenter/utils/helper"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func (h *CompanyHandler) GetCompanyById(w http.ResponseWriter, r *http.Request) {
	var (
		ctx       = context.TODO()
		vars      = mux.Vars(r)
		companyId = vars["company_id"]
		log       = logger.NewLogger(fmt.Sprintf("/v1/company/%s", companyId))
	)

	_, errToken := utils.ValidateTokenFromHeader(r)
	if errToken != nil {
		helper.ResponseErr(w, errToken, http.StatusUnauthorized)
		log.Error(errToken)
		return
	}

	company, jobs, err := h.UCCompany.GetCompanyById(ctx, companyId)
	if err != nil {
		helper.ResponseErr(w, err, http.StatusInternalServerError)
		log.Error(err)
		return
	}

	companyResponse := response.GetCompanyProfileResponse(company, jobs)
	helper.ResponseInterface(w, "success get profile company", companyResponse, http.StatusOK)
	log.InfoWithData("success Get profile compan", nil)
	return
}
