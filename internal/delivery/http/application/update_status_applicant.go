package application

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

func (h *ApplicationHandler) UpdateStatusApplicant(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = context.TODO()
		req     *request.RequestStatusApplicant
		decoder = json.NewDecoder(r.Body)
		log     = logger.NewLogger("/v1/admin/update/status-applicant")
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
		return
	}

	requestResult := request.NewApplicantRequest(req)

	err := h.UCApplication.UpdateStatusById(ctx, requestResult)
	if err != nil {
		helper.ResponseErr(w, err, http.StatusInternalServerError)
		log.Error(err)
		return
	}

	helper.Response(w, "success update applicant", http.StatusOK)
	log.InfoWithData("success update applicant", nil)
	return
}
