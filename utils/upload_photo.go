package utils

import (
	"CareerCenter/pkg/config"
	"CareerCenter/utils/exceptions"
	"io"
	"net/http"
	"os"
	"strings"
)

func UploadPhotoProfile(email string, r *http.Request, cfg config.Config) (string, error) {
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
	if fileType == "image/png" {
		header.Filename = "Photo" + name[0] + ".png"
	}
	if fileType == "image/jpg" {
		header.Filename = "Photo" + name[0] + ".jpg"
	}
	if fileType == "image/jpeg" {
		header.Filename = "Photo" + name[0] + ".jpeg"
	}

	path := cfg.PATH_IMAGE_UPLOAD + header.Filename

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		return "", err
	}

	return header.Filename, nil
}

func UploadImage(id string, attribut string, r *http.Request, cfg config.Config) (string, error) {
	_, err := ValitUuId(id)
	if err != nil {
		return "", err
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		return "", err
	}
	defer file.Close()

	fileType := header.Header.Get("Content-Type")
	if fileType != "image/png" && fileType != "image/jpg" && fileType != "image/jpeg" {
		return "", exceptions.ErrCustomString("Invalid file type, only PNG, JPG or JPEG is allowed")
	}

	if fileType == "image/png" {
		header.Filename = id + attribut + ".png"
	}
	if fileType == "image/jpg" {
		header.Filename = id + attribut + ".jpg"
	}
	if fileType == "image/jpeg" {
		header.Filename = id + attribut + ".jpeg"
	}
	path := cfg.PATH_IMAGE_UPLOAD + header.Filename

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		return "", err
	}

	return header.Filename, nil
}
