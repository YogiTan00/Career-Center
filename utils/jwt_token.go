package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

var sampleSecretKey = []byte("GoLinuxCloudKey")

func generateJWT(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(sampleSecretKey)

	if err != nil {
		fmt.Errorf("something went wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func validateToken(w http.ResponseWriter, r *http.Request) (err error) {

	if r.Header["Token"] == nil {
		fmt.Fprintf(w, "can not find token in header")
		return
	}

	token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error in parsing")
		}
		return sampleSecretKey, nil
	})

	if token == nil {
		fmt.Fprintf(w, "invalid token")
	}

	return nil
}
