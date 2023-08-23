package profile

import (
	"CareerCenter/internal/delivery/request"
	"CareerCenter/logger"
	"CareerCenter/utils"
	"CareerCenter/utils/helper"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func (h *ProfileHandler) UpdateEducation(w http.ResponseWriter, r *http.Request) {
	var (
		ctx         = context.TODO()
		req         *request.RequestEducation
		decoder     = json.NewDecoder(r.Body)
		vars        = mux.Vars(r)
		educationId = vars["education_id"]
		log         = logger.NewLogger(fmt.Sprintf("/v1/profile/update-education/%s", educationId))
	)

	errDecode := decoder.Decode(&req)
	if errDecode != nil {
		helper.ResponseErr(w, errDecode, http.StatusInternalServerError)
		log.Error(errDecode)
		return
	}

	_, errToken := utils.ValidateTokenFromHeader(r)
	if errToken != nil {
		helper.ResponseErr(w, errToken, http.StatusUnauthorized)
		log.Error(errToken)
		return
	}

	workExperience, err := request.NewUpdateEducation(req)
	if err != nil {
		helper.ResponseErr(w, err, http.StatusBadRequest)
		log.Error(err)
		return
	}

	err = h.UCProfile.UpdateEducation(ctx, educationId, workExperience)
	if err != nil {
		helper.ResponseErr(w, err, http.StatusInternalServerError)
		log.Error(err)
		return
	}

	helper.Response(w, "success update education", http.StatusOK)
	log.InfoWithData("success update education", nil)
	return
}
