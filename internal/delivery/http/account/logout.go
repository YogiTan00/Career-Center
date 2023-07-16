package account

import (
	"CareerCenter/logger"
	"CareerCenter/utils"
	"CareerCenter/utils/helper"
	"net/http"
)

func (h *AccountHandler) Logout(w http.ResponseWriter, r *http.Request) {
	var (
		log = logger.NewLogger("/v1/login")
	)
	_, errToken := utils.ValidateTokenFromHeader(r)
	if errToken != nil {
		helper.ResponseErr(w, errToken, http.StatusUnauthorized)
		return
	}

	utils.RemoveJwtInCookie(w)
	helper.Response(w, "success logout", http.StatusOK)
	log.InfoWithData("success logout", nil)
	return
}
