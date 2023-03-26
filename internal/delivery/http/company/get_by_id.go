package company

import (
	response2 "CareerCenter/internal/delivery/response"
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
		http.Error(w, errToken.Error(), http.StatusUnauthorized)
		return
	}

	company, jobs, err := h.UCCompany.GetCompanyById(ctx, companyId)
	if err != nil {
		response, errMap := response2.MapResponse(1, "cant get profile company")
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}
		w.Write(response)
	} else {
		response := response2.GetCompanyProfileResponse(company, jobs)
		result, errMap := response2.MapResponseInterface(0, "success Get profile company", response)
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}
		w.Write(result)
	}
}
