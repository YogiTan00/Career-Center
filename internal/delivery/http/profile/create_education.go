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
		log.Error(errDecode)
		return
	}

	user, errToken := utils.ValidateTokenFromHeader(r)
	if errToken != nil {
		helper.ResponseErr(w, errToken, http.StatusUnauthorized)
		log.Error(errToken)
		return
	}

	education, err := request.NewUpdateEducation(req)
	if err != nil {
		helper.ResponseErr(w, err, http.StatusBadRequest)
		return
	}

	err = h.UCProfile.CreateEducation(ctx, user.Email, education)
	if err != nil {
		helper.ResponseErr(w, err, http.StatusInternalServerError)
		log.Error(err)
		return
	}

	helper.Response(w, "success add education", http.StatusOK)
	log.InfoWithData("success add education", nil)
	return
}
