package account

import (
	"CareerCenter/domain/entity/account"
	"CareerCenter/internal/delivery/request"
	"CareerCenter/internal/delivery/response"
	"context"
	"encoding/json"
	"net/http"
)

func (h *AccountHandler) Login(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = context.TODO()
		req     request.RequestLogin
		decoder = json.NewDecoder(r.Body)
	)
	errDecode := decoder.Decode(&req)
	if errDecode != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error decode data"))
		return
	}
	buildLogin := &account.AccountDTO{
		Email:    req.Email,
		Password: req.Password,
	}

	token, err := h.UCAccount.Login(ctx, buildLogin.Email, buildLogin.Password)
	if err != nil {
		result, errMap := response.MapResponse(1, "wrong email or password")
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}
		w.Write(result)
	} else {
		http.SetCookie(w, token)
		result, errMap := response.MapResponseInterface(0, "success login", token.Value)
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}
		w.Write(result)
	}
}
