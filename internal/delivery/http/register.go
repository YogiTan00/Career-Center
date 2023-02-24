package http

import (
	"CareerCenter/domain/entity"
	"CareerCenter/internal/delivery/request"
	response2 "CareerCenter/internal/delivery/response"
	"context"
	"encoding/json"
	"net/http"
)

func (h *RegisterHandler) Register(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = context.TODO()
		req     request.RequestRegister
		decoder = json.NewDecoder(r.Body)
	)
	errDecode := decoder.Decode(&req)
	if errDecode != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error decode data"))
		return
	}

	buildRegister := &entity.AccountDTO{
		Email:    req.Email,
		Nama:     req.Nama,
		Password: req.Password,
	}

	errRegisterUseCase := h.UCRegister.Register(ctx, buildRegister)
	if errRegisterUseCase != nil {
		response, errMap := response2.MapResponse(1, "register error")
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	} else {
		response, errMap := response2.MapResponse(0, "Success")
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}
		w.Write(response)
	}
}
