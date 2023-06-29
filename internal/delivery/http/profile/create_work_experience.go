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

func (h *ProfileHandler) CreateWorkExperience(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = context.TODO()
		req     *request.RequestWorkExperience
		decoder = json.NewDecoder(r.Body)
		log     = logger.NewLogger(r.RequestURI)
	)
	errDecode := decoder.Decode(&req)
	if errDecode != nil {
		helper.ResponseErr(w, errDecode, http.StatusInternalServerError)
		log.General("", errDecode)
		return
	}

	user, errToken := utils.ValidateTokenFromHeader(r)
	if errToken != nil {
		helper.ResponseErr(w, errToken, http.StatusUnauthorized)
		log.General("", errToken)
		return
	}

	workExperience, err := request.NewUpdateWorkExperience(req)
	if err != nil {
		helper.ResponseErr(w, err, http.StatusBadRequest)
		log.General("", err)
		return
	}

	err = h.UCProfile.CreateWorkExperiencet(ctx, user.Email, workExperience)
	if err != nil {
		helper.ResponseErr(w, err, http.StatusInternalServerError)
		log.General("", err)
		return
	}

	helper.Response(w, "success add work experience", http.StatusInternalServerError)
	log.General("success add work experience", nil)
	return
}
