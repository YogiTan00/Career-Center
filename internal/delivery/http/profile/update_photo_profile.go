package profile

import (
	"CareerCenter/logger"
	"CareerCenter/utils"
	"CareerCenter/utils/helper"
	"context"
	"net/http"
)

func (h *ProfileHandler) UpdatePhotoProfile(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = context.TODO()
		log = logger.NewLogger("/v1/profile/update-photo")
	)

	user, errToken := utils.ValidateTokenFromHeader(r)
	if errToken != nil {
		helper.ResponseErr(w, errToken, http.StatusUnauthorized)
		log.Error(errToken)
		return
	}

	path, errUpload := utils.UploadPhoto(user.Email, r)
	if errUpload != nil {
		helper.ResponseErr(w, errUpload, http.StatusBadRequest)
		log.Error(errUpload)
		return
	}

	err := h.UCProfile.UpdatePhotoProfile(ctx, user.Email, path)
	if err != nil {
		helper.ResponseErr(w, err, http.StatusInternalServerError)
		log.Error(err)
		return
	}

	helper.Response(w, "success update photo profile", http.StatusOK)
	log.InfoWithData("success update photo profile", nil)
	return
}
