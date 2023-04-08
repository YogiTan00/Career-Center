package company

import (
	"CareerCenter/internal/delivery/response"
	"CareerCenter/utils"
	"CareerCenter/utils/helper"
	"context"
	"github.com/gorilla/mux"
	"net/http"
)

func (h *CompanyHandler) GetCompanyById(w http.ResponseWriter, r *http.Request) {
	var (
		ctx  = context.TODO()
		vars = mux.Vars(r)
	)

	companyId := vars["company_id"]

	_, errToken := utils.ValidateTokenFromHeader(r)
	if errToken != nil {
		helper.ResponseErr(w, errToken, http.StatusUnauthorized)
		return
	}

	company, jobs, err := h.UCCompany.GetCompanyById(ctx, companyId)
	if err != nil {
		helper.ResponseErr(w, err, http.StatusInternalServerError)
		return
	}

	companyResponse := response.GetCompanyProfileResponse(company, jobs)
	helper.ResponseInterface(w, "success Get profile company", companyResponse, http.StatusInternalServerError)
	return
}
