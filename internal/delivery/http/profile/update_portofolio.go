package profile

import (
	"CareerCenter/utils"
	"CareerCenter/utils/helper"
	"context"
	"net/http"
)

func (h *ProfileHandler) UpdatePortofolio(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = context.TODO()
	)

	email, errToken := utils.ValidateTokenFromHeader(r)
	if errToken != nil {
		helper.ResponseErr(w, errToken, http.StatusUnauthorized)
		return
	}

	path, errPdf := utils.UploadPDF(email, string(utils.TYPE_PORTOFOLIO), r)
	if errPdf != nil {
		helper.ResponseErr(w, errPdf, http.StatusBadRequest)
		return
	}

	err := h.UCProfile.UpdatePortofolio(ctx, email, path)
	if err != nil {
		helper.ResponseErr(w, err, http.StatusInternalServerError)
		return
	}

	helper.Response(w, "success update cv or resume", http.StatusInternalServerError)
	return
}
