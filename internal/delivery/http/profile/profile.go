package profile

import (
	"CareerCenter/internal/delivery/response"
	"CareerCenter/utils"
	"context"
	"net/http"
)

func (h *ProfileHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = context.TODO()
	)
	email, errToken := utils.ValidateTokenFromHeader(r)
	if errToken != nil {
		http.Error(w, errToken.Error(), http.StatusUnauthorized)
		return
	}

	profile, err := h.UCProfile.GetProfileByEmail(ctx, email)
	if err != nil {
		result, errMap := response.MapResponse(1, "cant get profile")
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}
		w.Write(result)
	} else {
		profileResponse := response.GetProfileResponse(profile)
		result, errMap := response.MapResponseInterface(0, "success get profile", profileResponse)
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}
		w.Write(result)
	}
}
