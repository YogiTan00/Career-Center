package account

import (
	"CareerCenter/internal/delivery/request"
	"CareerCenter/logger"
	"CareerCenter/utils/helper"
	"context"
	"encoding/json"
	"net/http"
)

func (h *AccountHandler) Login(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = context.TODO()
		req     *request.RequestLogin
		decoder = json.NewDecoder(r.Body)
		log     = logger.NewLogger("/v1/login")
	)

	errDecode := decoder.Decode(&req)
	if errDecode != nil {
		helper.ResponseErr(w, errDecode, http.StatusInternalServerError)
		log.General("", errDecode)
		return
	}

	buildLogin := request.NewLoginRequest(req)

	cookie, result, err := h.UCAccount.Login(ctx, buildLogin)
	if err != nil {
		helper.ResponseErr(w, err, http.StatusInternalServerError)
		log.General("", err)
		return
	}

	http.SetCookie(w, cookie)
	helper.ResponseInterface(w, "success login", result, http.StatusInternalServerError)
	log.General("success login", result)
	return
}
