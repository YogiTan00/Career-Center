package profile

import (
	"CareerCenter/logger"
	"CareerCenter/utils"
	"CareerCenter/utils/helper"
	"context"
	"net/http"
)

func (h *ProfileHandler) UpdatePortofolio(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = context.TODO()
		log = logger.NewLogger("/v1/profile/update-portofolio")
	)

	user, errToken := utils.ValidateTokenFromHeader(r)
	if errToken != nil {
		helper.ResponseErr(w, errToken, http.StatusUnauthorized)
		log.Error(errToken)
		return
	}

	path, errPdf := utils.UploadPDF(user.Email, string(utils.TYPE_PORTOFOLIO), r, h.cfg)
	if errPdf != nil {
		helper.ResponseErr(w, errPdf, http.StatusBadRequest)
		log.Error(errPdf)
		return
	}

	err := h.UCProfile.UpdatePortofolio(ctx, user.Email, path)
	if err != nil {
		helper.ResponseErr(w, err, http.StatusInternalServerError)
		log.Error(err)
		return
	}

	helper.Response(w, "success update cv or resume", http.StatusOK)
	log.InfoWithData("success update cv or resume", nil)
	return
}
