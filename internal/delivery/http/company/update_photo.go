package company

import (
	"CareerCenter/logger"
	"CareerCenter/utils"
	"CareerCenter/utils/exceptions"
	"CareerCenter/utils/helper"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func (h *CompanyHandler) UpdateLogoCompany(w http.ResponseWriter, r *http.Request) {
	var (
		ctx       = context.TODO()
		vars      = mux.Vars(r)
		companyId = vars["company_id"]
		log       = logger.NewLogger(fmt.Sprintf("/v1/admin/update-logo/%s", companyId))
	)

	user, errToken := utils.ValidateTokenFromHeader(r)
	if errToken != nil {
		helper.ResponseErr(w, errToken, http.StatusUnauthorized)
		log.Error(errToken)
		return
	}

	if !user.Admin() {
		helper.ResponseErr(w, exceptions.Unauthorized, http.StatusUnauthorized)
		log.Error(exceptions.Unauthorized)
		return
	}

	path, errUpload := utils.UploadImage(companyId, "Logo", r, h.cfg)
	if errUpload != nil {
		helper.ResponseErr(w, errUpload, http.StatusBadRequest)
		log.Error(errUpload)
		return
	}

	err := h.UCCompany.UpdateLogoCompany(ctx, companyId, path)
	if err != nil {
		helper.ResponseErr(w, err, http.StatusInternalServerError)
		log.Error(err)
		return
	}

	helper.Response(w, "success update photo profile", http.StatusOK)
	log.InfoWithData("success update photo profile", nil)
	return
}
