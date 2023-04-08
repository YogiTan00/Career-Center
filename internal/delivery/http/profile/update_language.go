package profile

import (
	"CareerCenter/internal/delivery/request"
	"CareerCenter/internal/delivery/response"
	"CareerCenter/utils"
	"CareerCenter/utils/helper"
	"context"
	"encoding/json"
	"net/http"
)

func (h *ProfileHandler) UpdateLanguage(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = context.TODO()
		req     *request.RequestLanguage
		decoder = json.NewDecoder(r.Body)
	)

	errDecode := decoder.Decode(&req)
	if errDecode != nil {
		helper.ResponseErr(w, errDecode, http.StatusInternalServerError)
		return
	}

	email, errToken := utils.ValidateTokenFromHeader(r)
	if errToken != nil {
		helper.ResponseErr(w, errToken, http.StatusUnauthorized)
		return
	}

	err := h.UCProfile.UpdateLanguage(ctx, email, req.Language)
	if err != nil {
		helper.ResponseErr(w, err, http.StatusInternalServerError)
		return
	}

	result, errMap := response.MapResponse(0, "success update language")
	if errMap != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error mapping data"))
	}
	w.Write(result)
	return
}
