package handler

import (
	"CareerCenter/logger"
	"fmt"
	"net/http"
)

func ParamHandlerWithoutInput(w http.ResponseWriter, r *http.Request) {
	var (
		log = logger.NewLogger(r.RequestURI)
	)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "SUCCESS OK")
	log.General("SUCCESS OK", nil)
}

func GetImage(w http.ResponseWriter, r *http.Request) {
	var (
		log = logger.NewLogger(r.RequestURI)
	)
	filename := "uploads/image/" + r.URL.Query().Get("filename")
	http.ServeFile(w, r, filename)
	log.General("get image", filename)
}

func GetPdf(w http.ResponseWriter, r *http.Request) {
	var (
		log = logger.NewLogger(r.RequestURI)
	)
	filename := "uploads/pdf/" + r.URL.Query().Get("filename")
	http.ServeFile(w, r, filename)
	log.General("get pdf", filename)
}
