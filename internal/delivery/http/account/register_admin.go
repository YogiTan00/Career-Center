package account

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

func (h *AccountHandler) RegisterAdmin(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = context.TODO()
		req     *request.RequestRegister
		decoder = json.NewDecoder(r.Body)
		log     = logger.NewLogger("/v1/admin/register")
	)
	errDecode := decoder.Decode(&req)
	if errDecode != nil {
		helper.ResponseErr(w, errDecode, http.StatusInternalServerError)
		log.General("", errDecode)
		return
	}

	user, errToken := utils.ValidateTokenFromHeader(r)
	if errToken != nil {
		helper.ResponseErr(w, errToken, http.StatusUnauthorized)
		log.General("", errToken)
		return
	}

	if !user.Admin() {
		helper.ResponseErr(w, exceptions.Unauthorized, http.StatusUnauthorized)
		log.General("", exceptions.Unauthorized)
		return
	}

	buildRegister := request.NewRegisterByAdminRequest(req)

	errRegisterUseCase := h.UCAccount.RegisterAdmin(ctx, user.Role, buildRegister)
	if errRegisterUseCase != nil {
		helper.ResponseErr(w, errRegisterUseCase, http.StatusInternalServerError)
		log.General("", errRegisterUseCase)
		return
	}

	helper.Response(w, "success register", http.StatusInternalServerError)
	log.General("Success register", errRegisterUseCase)
	return
}
