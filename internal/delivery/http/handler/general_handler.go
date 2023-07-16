package handler

import (
	"CareerCenter/logger"
	"fmt"
	"net/http"
)

func ParamHandlerWithoutInput(w http.ResponseWriter, r *http.Request) {
	var (
		log = logger.NewLogger("/")
	)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "SUCCESS OK")
	log.InfoWithData("SUCCESS OK", nil)
}

func GetImage(w http.ResponseWriter, r *http.Request) {
	var (
		log = logger.NewLogger("/v1/photo")
	)
	filename := "uploads/image/" + r.URL.Query().Get("filename")
	http.ServeFile(w, r, filename)
	w.WriteHeader(http.StatusOK)
	log.InfoWithData("get image", filename)
}

func GetPdf(w http.ResponseWriter, r *http.Request) {
	var (
		log = logger.NewLogger("/v1/pdf")
	)
	filename := "uploads/pdf/" + r.URL.Query().Get("filename")
	http.ServeFile(w, r, filename)
	w.WriteHeader(http.StatusOK)
	log.InfoWithData("get pdf", filename)
}
