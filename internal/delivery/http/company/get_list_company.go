package company

import (
	"CareerCenter/internal/delivery/request"
	response2 "CareerCenter/internal/delivery/response"
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
		response, errMap := response2.MapResponse(1, errFilter.Error())
		if errMap != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error mapping data"))
		}
		w.Write(response)
		return
	}

	company, err := h.UCCompany.GetListCompany(ctx, filter)
	if err != nil {
		response, errMap := response2.MapResponse(1, "cant get list company")
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}
		w.Write(response)
	} else {
		response := response2.GetListCompanyResponse(company)
		result, errMap := response2.MapResponseInterface(0, "success Get list company", response)
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}
		w.Write(result)
	}
}
