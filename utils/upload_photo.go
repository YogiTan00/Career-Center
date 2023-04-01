package utils

import (
	"CareerCenter/utils/exceptions"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func UploadPhoto(r *http.Request) (string, error) {
	file, header, err := r.FormFile("photo")
	if err != nil {
		return "", err
	}
	defer file.Close()

	fileType := header.Header.Get("Content-Type")
	if fileType != "image/png" {
		return "", exceptions.ErrCustomString("Invalid file type, only PNG is allowed")
	}

	timeNow := time.Now().Format(fmt.Sprintf("2006%s01%s02%s15%s04%s05", RandomString(2), RandomString(2), RandomString(2), RandomString(2), RandomString(2)))
	header.Filename = timeNow + ".png"

	path := "uploads/" + header.Filename
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("http://%s/photo/%s", r.Host, header.Filename)

	return url, nil
}
