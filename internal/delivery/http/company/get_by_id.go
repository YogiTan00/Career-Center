package company

import (
	"CareerCenter/internal/delivery/response"
	"CareerCenter/utils"
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
		result, errMap := response.MapResponse(1, errToken.Error())
		if errMap != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Error mapping data"))
		}
		w.Write(result)
		return
	}

	company, jobs, err := h.UCCompany.GetCompanyById(ctx, companyId)
	if err != nil {
		result, errMap := response.MapResponse(1, "cant get profile company")
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}
		w.Write(result)
		return
	}

	companyResponse := response.GetCompanyProfileResponse(company, jobs)
	result, errMap := response.MapResponseInterface(0, "success Get profile company", companyResponse)
	if errMap != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error mapping data"))
	}
	w.Write(result)
	return
}
