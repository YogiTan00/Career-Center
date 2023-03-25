package profile

import (
	response2 "CareerCenter/internal/delivery/response"
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
		response, errMap := response2.MapResponse(1, "cant get profile")
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}
		w.Write(response)
	} else {
		response := response2.GetProfileResponse(profile)
		result, errMap := response2.MapResponseInterface(0, "success get profile", response)
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}
		w.Write(result)
	}
}
