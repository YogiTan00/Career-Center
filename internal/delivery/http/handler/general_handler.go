package handler

import (
	"fmt"
	"net/http"
)

func ParamHandlerWithoutInput(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "SUCCESS OK")
}

func GetImage(w http.ResponseWriter, r *http.Request) {
	filename := "uploads/image/" + r.URL.Query().Get("filename")
	http.ServeFile(w, r, filename)
}

func GetPdf(w http.ResponseWriter, r *http.Request) {
	filename := "uploads/pdf/" + r.URL.Query().Get("filename")
	http.ServeFile(w, r, filename)
}
