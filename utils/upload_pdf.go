package utils

import (
	"CareerCenter/utils/exceptions"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type TypePdf string

const (
	TYPE_CV_RESUME  TypePdf = "CvResume"
	TYPE_PORTOFOLIO TypePdf = "Portofolio"
)

func UploadPDF(email string, typePdf string, r *http.Request) (string, error) {
	// Menerima file dari form
	file, header, err := r.FormFile("file")
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Mendapatkan ekstensi file
	ext := filepath.Ext(header.Filename)

	if ext != ".pdf" {
		return "", exceptions.ErrCustomString("wrong format file only pdf")
	}

	// Menyimpan file di folder "uploads"
	name := strings.Split(email, "@")
	header.Filename = name[0] + typePdf + ext
	path := filepath.Join("uploads/pdf", header.Filename)
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("http://%s/pdf/%s", r.Host, header.Filename)

	return url, nil
}
