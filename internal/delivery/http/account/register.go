package account

import (
	"CareerCenter/internal/delivery/request"
	"CareerCenter/utils/helper"
	"context"
	"encoding/json"
	"net/http"
)

func (h *AccountHandler) Register(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = context.TODO()
		req     *request.RequestRegister
		decoder = json.NewDecoder(r.Body)
	)
	errDecode := decoder.Decode(&req)
	if errDecode != nil {
		helper.ResponseErr(w, errDecode, http.StatusInternalServerError)
		return
	}

	buildRegister := request.NewRegisterRequest(req)

	errRegisterUseCase := h.UCAccount.Register(ctx, buildRegister)
	if errRegisterUseCase != nil {
		helper.ResponseErr(w, errRegisterUseCase, http.StatusInternalServerError)
		return
	}

	helper.Response(w, "Success register", http.StatusInternalServerError)
	return
}
