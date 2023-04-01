package account

import (
	"CareerCenter/internal/delivery/request"
	"CareerCenter/internal/delivery/response"
	"CareerCenter/utils"
	"context"
	"encoding/json"
	"net/http"
)

func (h *AccountHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = context.TODO()
		req     *request.RequestPassword
		decoder = json.NewDecoder(r.Body)
	)
	errDecode := decoder.Decode(&req)
	if errDecode != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error decode data"))
		return
	}

	user, errToken := utils.ValidateTokenFromHeader(r)
	if errToken != nil {
		http.Error(w, errToken.Error(), http.StatusUnauthorized)
		return
	}

	password := request.NewPassword(req)

	err := h.UCAccount.UpdatePassword(ctx, user, password)
	if err != nil {
		response, errMap := response.MapResponse(1, err.Error())
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}
		w.Write(response)
		return
	}

	result, errMap := response.MapResponse(0, "success change password")
	if errMap != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error mapping data"))
	}
	w.Write(result)
	return
}
