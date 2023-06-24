package account

import (
	"CareerCenter/internal/delivery/request"
	"CareerCenter/logger"
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
		log     = logger.NewLogger("/v1/register")
	)
	errDecode := decoder.Decode(&req)
	if errDecode != nil {
		helper.ResponseErr(w, errDecode, http.StatusInternalServerError)
		log.General("", errDecode)
		return
	}

	buildRegister := request.NewRegisterRequest(req)

	errRegisterUseCase := h.UCAccount.Register(ctx, buildRegister)
	if errRegisterUseCase != nil {
		helper.ResponseErr(w, errRegisterUseCase, http.StatusInternalServerError)
		log.General("", errRegisterUseCase)
		return
	}

	helper.Response(w, "success register", http.StatusInternalServerError)
	log.General("Success register", errRegisterUseCase)
	return
}
