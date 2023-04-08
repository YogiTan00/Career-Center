package account

import (
	"CareerCenter/internal/delivery/request"
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
	)
	errDecode := decoder.Decode(&req)
	if errDecode != nil {
		helper.ResponseErr(w, errDecode, http.StatusInternalServerError)
		return
	}

	buildLogin := request.NewLoginRequest(req)

	token, err := h.UCAccount.Login(ctx, buildLogin.Email, buildLogin.Password)
	if err != nil {
		helper.ResponseErr(w, err, http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, token)
	helper.ResponseInterface(w, "success login", token.Value, http.StatusInternalServerError)
	return
}
