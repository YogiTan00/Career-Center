package account

import (
	"CareerCenter/internal/delivery/request"
	"CareerCenter/logger"
	"CareerCenter/utils/helper"
	"context"
	"encoding/json"
	"net/http"
)

func (h *AccountHandler) ForgetPassword(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = context.TODO()
		req     *request.RequestForgetPassword
		decoder = json.NewDecoder(r.Body)
		log     = logger.NewLogger("/v1/forget-password")
	)
	errDecode := decoder.Decode(&req)
	if errDecode != nil {
		helper.ResponseErr(w, errDecode, http.StatusInternalServerError)
		log.Error(errDecode)
		return
	}

	err := h.UCAccount.ForgetPassword(ctx, req.Email)
	if err != nil {
		helper.ResponseErr(w, errDecode, http.StatusBadRequest)
		log.Error(err)
		return
	}

	helper.Response(w, "success send forget password", http.StatusOK)
	log.InfoWithData("success send forget password", nil)
	return
}
