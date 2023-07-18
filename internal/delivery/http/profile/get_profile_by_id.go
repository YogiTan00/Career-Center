package profile

import (
	"CareerCenter/internal/delivery/response"
	"CareerCenter/logger"
	"CareerCenter/utils"
	"CareerCenter/utils/helper"
	"context"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *ProfileHandler) GetProfileByEmail(w http.ResponseWriter, r *http.Request) {
	var (
		ctx   = context.TODO()
		vars  = mux.Vars(r)
		email = vars["email"]
		log   = logger.NewLogger("/v1/profile/{email}")
	)
	_, errToken := utils.ValidateTokenFromHeader(r)
	if errToken != nil {
		helper.ResponseErr(w, errToken, http.StatusUnauthorized)
		log.Error(errToken)
		return
	}

	profile, err := h.UCProfile.GetProfileByEmail(ctx, email)
	if profile == nil {
		helper.ResponseErrWithData(w, errors.New("data not found"), http.StatusBadRequest, nil)
		log.Error(err)
		return
	}
	if err != nil {
		helper.ResponseErr(w, err, http.StatusInternalServerError)
		log.Error(err)
		return
	}

	profileResponse := response.GetProfileResponse(profile, h.cfg)
	helper.ResponseInterface(w, "success get profile by email", profileResponse, http.StatusOK)
	log.InfoWithData("success get profile by email", profileResponse)
	return
}
