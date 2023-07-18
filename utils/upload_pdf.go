package utils

import (
	"CareerCenter/pkg/config"
	"CareerCenter/utils/exceptions"
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

func UploadPDF(email string, typePdf string, r *http.Request, cfg config.Config) (string, error) {
	// Menerima file dari form
	file, header, err := r.FormFile("file")
	if err != nil {
		return "", err
	}
	defer file.Close()
	// Destination
	ext := filepath.Ext(header.Filename)
	// Mendapatkan ekstensi file
	if ext != ".pdf" {
		return "", exceptions.ErrCustomString("wrong format file only pdf")
	}
	// Menyimpan file di folder "uploads"
	name := strings.Split(email, "@")
	filename := name[0] + typePdf + ext

	path := cfg.PATH_FILE_UPLOAD + filename

	pathMeta := cfg.PATH_FILE_UPLOAD_META + filename
	dst, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer dst.Close()
	// Copy
	if _, err = io.Copy(dst, file); err != nil {
		return "", err
	}

	return pathMeta, nil
}
