package profile

import (
	"CareerCenter/internal/delivery/response"
	"CareerCenter/logger"
	"CareerCenter/utils"
	"CareerCenter/utils/helper"
	"context"
	"net/http"
)

func (h *ProfileHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = context.TODO()
		log = logger.NewLogger("/v1/profile")
	)
	user, errToken := utils.ValidateTokenFromHeader(r)
	if errToken != nil {
		helper.ResponseErr(w, errToken, http.StatusUnauthorized)
		log.Error(errToken)
		return
	}

	profile, err := h.UCProfile.GetProfileByEmail(ctx, user.Email)
	if err != nil {
		helper.ResponseErr(w, err, http.StatusInternalServerError)
		log.Error(err)
		return
	}

	profileResponse := response.GetProfileResponse(profile, h.cfg)
	helper.ResponseInterface(w, "success get profile", profileResponse, http.StatusOK)
	log.InfoWithData("success get profile", profileResponse)
	return
}
