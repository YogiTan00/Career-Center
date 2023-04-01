package company

import (
	"CareerCenter/internal/delivery/request"
	"CareerCenter/internal/delivery/response"
	"CareerCenter/utils"
	"context"
	"net/http"
)

func (h *CompanyHandler) GetListCompany(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = context.TODO()
		req request.RequestFilter
	)

	_, errToken := utils.ValidateTokenFromHeader(r)
	if errToken != nil {
		http.Error(w, errToken.Error(), http.StatusUnauthorized)
		return
	}

	filter, errFilter := request.FilterGeneral(r, &req)
	if errFilter != nil {
		result, errMap := response.MapResponse(1, errFilter.Error())
		if errMap != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error mapping data"))
		}
		w.Write(result)
		return
	}

	company, err := h.UCCompany.GetListCompany(ctx, filter)
	if err != nil {
		result, errMap := response.MapResponse(1, "cant get list company")
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}
		w.Write(result)
		return
	}

	companyResponse := response.GetListCompanyResponse(company)
	result, errMap := response.MapResponseInterface(0, "success Get list company", companyResponse)
	if errMap != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error mapping data"))
	}
	w.Write(result)
	return
}
