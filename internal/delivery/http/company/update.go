package company

import (
	"CareerCenter/internal/delivery/request"
	"CareerCenter/logger"
	"CareerCenter/utils"
	"CareerCenter/utils/exceptions"
	"CareerCenter/utils/helper"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func (h *CompanyHandler) UpdateCompany(w http.ResponseWriter, r *http.Request) {
	var (
		ctx       = context.TODO()
		req       *request.RequestCompany
		decoder   = json.NewDecoder(r.Body)
		vars      = mux.Vars(r)
		companyId = vars["company_id"]
		log       = logger.NewLogger(fmt.Sprintf("/v1/admin/update-company/%s", companyId))
	)

	errDecode := decoder.Decode(&req)
	if errDecode != nil {
		helper.ResponseErr(w, errDecode, http.StatusInternalServerError)
		log.Error(errDecode)
		return
	}

	user, errToken := utils.ValidateTokenFromHeader(r)
	if errToken != nil {
		helper.ResponseErr(w, errToken, http.StatusUnauthorized)
		log.Error(errToken)
		return
	}

	if !user.Admin() {
		helper.ResponseErr(w, exceptions.Unauthorized, http.StatusUnauthorized)
		log.Error(exceptions.Unauthorized)
		return
	}

	company := request.NewCompanyRequest(req)

	err := h.UCCompany.UpdateCompanyById(ctx, companyId, company)
	if err != nil {
		helper.ResponseErr(w, err, http.StatusInternalServerError)
		log.Error(err)
		return
	}

	helper.Response(w, "success update company", http.StatusOK)
	log.InfoWithData("success update company", nil)
	return
}
