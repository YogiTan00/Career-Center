package profile

import (
	"CareerCenter/utils"
	"CareerCenter/utils/helper"
	"context"
	"net/http"
)

func (h *ProfileHandler) UpdatePhotoProfile(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = context.TODO()
	)

	email, errToken := utils.ValidateTokenFromHeader(r)
	if errToken != nil {
		helper.ResponseErr(w, errToken, http.StatusUnauthorized)
		return
	}

	path, errUpload := utils.UploadPhoto(r)
	if errUpload != nil {
		helper.ResponseErr(w, errUpload, http.StatusInternalServerError)
		return
	}

	err := h.UCProfile.UpdatePhotoProfile(ctx, email, path)
	if err != nil {
		helper.ResponseErr(w, err, http.StatusInternalServerError)
		return
	}

	helper.Response(w, "success update photo profile", http.StatusInternalServerError)
	return
}
