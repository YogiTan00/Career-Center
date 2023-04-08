package profile

import (
	"CareerCenter/internal/delivery/response"
	"CareerCenter/utils"
	"CareerCenter/utils/helper"
	"context"
	"net/http"
)

func (h *ProfileHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = context.TODO()
	)
	email, errToken := utils.ValidateTokenFromHeader(r)
	if errToken != nil {
		helper.ResponseErr(w, errToken, http.StatusUnauthorized)
		return
	}

	profile, err := h.UCProfile.GetProfileByEmail(ctx, email)
	if err != nil {
		helper.ResponseErr(w, err, http.StatusInternalServerError)
		return
	}

	profileResponse := response.GetProfileResponse(profile)
	helper.ResponseInterface(w, "success get profile", profileResponse, http.StatusInternalServerError)
	return
}
