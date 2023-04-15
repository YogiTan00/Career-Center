package account

import (
	"CareerCenter/utils"
	"CareerCenter/utils/helper"
	"net/http"
)

func (h *AccountHandler) Logout(w http.ResponseWriter, r *http.Request) {
	_, errToken := utils.ValidateTokenFromHeader(r)
	if errToken != nil {
		helper.ResponseErr(w, errToken, http.StatusUnauthorized)
		return
	}

	utils.RemoveJwtInCookie(w)
	helper.Response(w, "success logout", http.StatusInternalServerError)
	return
}
