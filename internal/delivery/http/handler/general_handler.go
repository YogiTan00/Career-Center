package handler

import (
	"CareerCenter/logger"
	"CareerCenter/pkg/config"
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
		cfg config.Config
		log = logger.NewLogger("/v1/photo")
	)
	filename := cfg.PATH_IMAGE_UPLOAD + r.URL.Query().Get("filename")
	http.ServeFile(w, r, filename)
	w.WriteHeader(http.StatusOK)
	log.InfoWithData("get image", filename)
}

func GetPdf(w http.ResponseWriter, r *http.Request) {
	var (
		cfg config.Config
		log = logger.NewLogger("/v1/pdf")
	)
	filename := cfg.PATH_FILE_UPLOAD + r.URL.Query().Get("filename")
	http.ServeFile(w, r, filename)
	w.WriteHeader(http.StatusOK)
	log.InfoWithData("get pdf", filename)
}
