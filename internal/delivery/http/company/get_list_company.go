package company

import (
	"CareerCenter/internal/delivery/request"
	"CareerCenter/internal/delivery/response"
	"CareerCenter/logger"
	"CareerCenter/utils"
	"CareerCenter/utils/helper"
	"context"
	"net/http"
)

func (h *CompanyHandler) GetListCompany(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = context.TODO()
		req request.RequestFilter
		log = logger.NewLogger("/v1/list-company")
	)

	_, errToken := utils.ValidateTokenFromHeader(r)
	if errToken != nil {
		helper.ResponseErr(w, errToken, http.StatusUnauthorized)
		log.General("", errToken)
		return
	}

	filter, errFilter := request.FilterGeneral(r, &req)
	if errFilter != nil {
		helper.ResponseErr(w, errFilter, http.StatusBadRequest)
		log.General("", errFilter)
		return
	}

	company, err := h.UCCompany.GetListCompany(ctx, filter)
	if err != nil {
		helper.ResponseErr(w, err, http.StatusInternalServerError)
		log.General("", err)
		return
	}

	companyResponse := response.GetListCompanyResponse(company)
	helper.ResponseInterface(w, "success get list company", companyResponse, http.StatusOK)
	log.General("success login", companyResponse)
	return
}
