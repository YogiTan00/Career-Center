package utils

import (
	"net/http"
)

func UploadPDF(w http.ResponseWriter, r *http.Request) {
	//file, header, err := r.FormFile("file")
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusBadRequest)
	//	return
	//}
	//defer file.Close()
	//
	//data, err := io.ReadAll(file)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
}
