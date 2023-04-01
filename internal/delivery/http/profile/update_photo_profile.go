package profile

import (
	"CareerCenter/internal/delivery/response"
	"CareerCenter/utils"
	"context"
	"net/http"
)

func (h *ProfileHandler) UpdatePhotoProfile(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = context.TODO()
	)

	email, errToken := utils.ValidateTokenFromHeader(r)
	if errToken != nil {
		http.Error(w, errToken.Error(), http.StatusUnauthorized)
		return
	}

	path, errUpload := utils.UploadPhoto(r)
	if errUpload != nil {
		result, errMap := response.MapResponse(1, errUpload.Error())
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}
		w.Write(result)
		return
	}

	err := h.UCProfile.UpdatePhotoProfile(ctx, email, path)
	if err != nil {
		result, errMap := response.MapResponse(1, err.Error())
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}
		w.Write(result)
		return
	}

	result, errMap := response.MapResponse(0, "success update photo profile")
	if errMap != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error mapping data"))
	}
	w.Write(result)
	return
}
