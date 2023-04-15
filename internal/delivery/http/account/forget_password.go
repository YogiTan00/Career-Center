package account

import (
	"CareerCenter/internal/delivery/request"
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
	)
	errDecode := decoder.Decode(&req)
	if errDecode != nil {
		helper.ResponseErr(w, errDecode, http.StatusInternalServerError)
		return
	}

	err := h.UCAccount.ForgetPassword(ctx, req.Email)
	if err != nil {
		helper.ResponseErr(w, errDecode, http.StatusBadRequest)
		return
	}

	helper.Response(w, "success send forget password", http.StatusOK)
	return
}
