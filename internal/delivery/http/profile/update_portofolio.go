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
		log = logger.NewLogger(r.RequestURI)
	)

	user, errToken := utils.ValidateTokenFromHeader(r)
	if errToken != nil {
		helper.ResponseErr(w, errToken, http.StatusUnauthorized)
		log.General("", errToken)
		return
	}

	path, errPdf := utils.UploadPDF(user.Email, string(utils.TYPE_PORTOFOLIO), r)
	if errPdf != nil {
		helper.ResponseErr(w, errPdf, http.StatusBadRequest)
		log.General("", errPdf)
		return
	}

	err := h.UCProfile.UpdatePortofolio(ctx, user.Email, path)
	if err != nil {
		helper.ResponseErr(w, err, http.StatusInternalServerError)
		log.General("", err)
		return
	}

	helper.Response(w, "success update cv or resume", http.StatusInternalServerError)
	log.General("success update cv or resume", nil)
	return
}
