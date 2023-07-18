package utils

import (
	"CareerCenter/pkg/config"
	"CareerCenter/utils/exceptions"
	"io"
	"net/http"
	"os"
	"strings"
)

func UploadPhoto(email string, r *http.Request, cfg config.Config) (string, error) {
	file, header, err := r.FormFile("photo")
	if err != nil {
		return "", err
	}
	defer file.Close()

	fileType := header.Header.Get("Content-Type")
	if fileType != "image/png" && fileType != "image/jpg" && fileType != "image/jpeg" {
		return "", exceptions.ErrCustomString("Invalid file type, only PNG, JPG or JPEG is allowed")
	}

	name := strings.Split(email, "@")
	header.Filename = name[0] + "Photo" + ".png"
	path := cfg.PATH_IMAGE_UPLOAD + header.Filename
	pathMeta := cfg.PATH_IMAGE_UPLOAD_META + header.Filename

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		return "", err
	}

	return pathMeta, nil
}
