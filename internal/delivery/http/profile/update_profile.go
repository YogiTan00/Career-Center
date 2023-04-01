package profile

import (
	"CareerCenter/internal/delivery/request"
	"CareerCenter/internal/delivery/response"
	"CareerCenter/utils"
	"context"
	"encoding/json"
	"net/http"
)

func (h *ProfileHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = context.TODO()
		req     *request.RequestUpdateProfile
		decoder = json.NewDecoder(r.Body)
	)
	errDecode := decoder.Decode(&req)
	if errDecode != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error decode data"))
		return
	}

	email, errToken := utils.ValidateTokenFromHeader(r)
	if errToken != nil {
		http.Error(w, errToken.Error(), http.StatusUnauthorized)
		return
	}

	buildProfile := request.NewUpdateProfileRequest(req)

	err := h.UCProfile.UpdateProfile(ctx, email, buildProfile)
	if err != nil {
		result, errMap := response.MapResponse(1, err.Error())
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}
		w.Write(result)
		return
	}

	result, errMap := response.MapResponse(0, "success update profile")
	if errMap != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error mapping data"))
	}
	w.Write(result)
	return
}
