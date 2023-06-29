package account

import (
	"CareerCenter/logger"
	"CareerCenter/utils"
	"CareerCenter/utils/helper"
	"net/http"
)

func (h *AccountHandler) Logout(w http.ResponseWriter, r *http.Request) {
	var (
		log = logger.NewLogger(r.RequestURI)
	)
	_, errToken := utils.ValidateTokenFromHeader(r)
	if errToken != nil {
		helper.ResponseErr(w, errToken, http.StatusUnauthorized)
		return
	}

	utils.RemoveJwtInCookie(w)
	helper.Response(w, "success logout", http.StatusInternalServerError)
	log.General("success logout", nil)
	return
}
