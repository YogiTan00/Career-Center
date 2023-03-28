package application

import (
	"CareerCenter/internal/delivery/request"
	response2 "CareerCenter/internal/delivery/response"
	"CareerCenter/utils"
	"context"
	"encoding/json"
	"net/http"
)

func (h *ApplicationHandler) SendApplication(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = context.TODO()
		req     *request.RequestApplication
		decoder = json.NewDecoder(r.Body)
	)
	user, errToken := utils.ValidateTokenFromHeader(r)
	if errToken != nil {
		http.Error(w, errToken.Error(), http.StatusUnauthorized)
		return
	}

	errDecode := decoder.Decode(&req)
	if errDecode != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error decode data"))
		return
	}

	err := h.UCApplication.SendApplication(ctx, user, req.CompanyId)
	if err != nil {
		response, errMap := response2.MapResponse(1, err.Error())
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}
		w.Write(response)
	} else {
		result, errMap := response2.MapResponse(0, "success send application")
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}
		w.Write(result)
	}
}