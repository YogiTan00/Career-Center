package company

import (
	"CareerCenter/internal/delivery/request"
	"CareerCenter/logger"
	"CareerCenter/utils"
	"CareerCenter/utils/exceptions"
	"CareerCenter/utils/helper"
	"context"
	"encoding/json"
	"net/http"
)

func (h *CompanyHandler) CreatCompany(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = context.TODO()
		req     *request.RequestCompany
		decoder = json.NewDecoder(r.Body)
		log     = logger.NewLogger("/v1/admin/company")
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

	buildCreateCompany := request.NewCompanyRequest(req)

	errCreateUseCase := h.UCCompany.CreateCompany(ctx, buildCreateCompany)
	if errCreateUseCase != nil {
		helper.ResponseErr(w, errCreateUseCase, http.StatusInternalServerError)
		log.Error(errCreateUseCase)
		return
	}

	helper.Response(w, "success create company", http.StatusOK)
	log.InfoWithData("Success create company", errCreateUseCase)
	return
}
