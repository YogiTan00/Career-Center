package account

import (
	"CareerCenter/domain/entity"
	"CareerCenter/internal/delivery/request"
	response2 "CareerCenter/internal/delivery/response"
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
	buildLogin := &entity.AccountDTO{
		Email:    req.Email,
		Password: req.Password,
	}

	token, err := h.UCAccount.Login(ctx, buildLogin.Email, buildLogin.Password)
	if err != nil {
		response, errMap := response2.MapResponse(1, "wrong email or password")
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}
		w.Write(response)
	} else {
		http.SetCookie(w, token)
		response, errMap := response2.MapResponseLogin(0, "success login", token.Value)
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}
		w.Write(response)
	}
}
