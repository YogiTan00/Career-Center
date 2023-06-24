package profile

import (
	"CareerCenter/internal/delivery/request"
	"CareerCenter/logger"
	"CareerCenter/utils"
	"CareerCenter/utils/helper"
	"context"
	"encoding/json"
	"net/http"
)

func (h *ProfileHandler) CreateEducation(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = context.TODO()
		req     *request.RequestEducation
		decoder = json.NewDecoder(r.Body)
		log     = logger.NewLogger("/v1/profile/add-education")
	)
	errDecode := decoder.Decode(&req)
	if errDecode != nil {
		helper.ResponseErr(w, errDecode, http.StatusInternalServerError)
		log.General("", errDecode)
		return
	}

	email, errToken := utils.ValidateTokenFromHeader(r)
	if errToken != nil {
		helper.ResponseErr(w, errToken, http.StatusUnauthorized)
		log.General("", errToken)
		return
	}

	education, err := request.NewUpdateEducation(req)
	if err != nil {
		helper.ResponseErr(w, err, http.StatusBadRequest)
		return
	}

	err = h.UCProfile.CreateEducation(ctx, email, education)
	if err != nil {
		helper.ResponseErr(w, err, http.StatusInternalServerError)
		log.General("", err)
		return
	}

	helper.Response(w, "success add education", http.StatusInternalServerError)
	log.General("success add education", nil)
	return
}
