package account

import (
	"CareerCenter/internal/delivery/request"
	"CareerCenter/logger"
	"CareerCenter/utils/helper"
	"context"
	"encoding/json"
	"net/http"
)

func (h *AccountHandler) ForgetPasswordUpdate(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = context.TODO()
		req     *request.RequestForgetPasswordUpdate
		decoder = json.NewDecoder(r.Body)
		log     = logger.NewLogger("/v1/forget-password/update")
	)
	errDecode := decoder.Decode(&req)
	if errDecode != nil {
		helper.ResponseErr(w, errDecode, http.StatusInternalServerError)
		log.Error(errDecode)
		return
	}

	err := h.UCAccount.ForgetPasswordUpdate(ctx, req)
	if err != nil {
		helper.ResponseErr(w, err, http.StatusBadRequest)
		log.Error(err)
		return
	}

	helper.Response(w, "success update forget password", http.StatusOK)
	log.InfoWithData("success update forget password", nil)
	return
}
