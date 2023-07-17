package account

import (
	"CareerCenter/internal/delivery/request"
	"CareerCenter/logger"
	"CareerCenter/utils/helper"
	"context"
	"encoding/json"
	"net/http"
)

func (h *AccountHandler) SubmitOtp(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = context.TODO()
		req     *request.RequestSubmitOtp
		decoder = json.NewDecoder(r.Body)
		log     = logger.NewLogger("/v1/forget-password")
	)
	errDecode := decoder.Decode(&req)
	if errDecode != nil {
		helper.ResponseErr(w, errDecode, http.StatusInternalServerError)
		log.Error(errDecode)
		return
	}

	err := h.UCAccount.SubmitOtp(ctx, req.Email, req.Otp)
	if err != nil {
		helper.ResponseErr(w, err, http.StatusBadRequest)
		log.Error(err)
		return
	}

	helper.Response(w, "success submit code verification", http.StatusOK)
	log.InfoWithData("success submit code verification", nil)
	return
}
