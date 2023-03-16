package response

import (
	"encoding/json"
)

type CustomReponseSingleAuth struct {
	Status *Status
	Data   string
}

func MapResponseLogin(code int, message string, token string) ([]byte, error) {
	httpResponse := &CustomReponseSingleAuth{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: token,
	}

	respJson, err := json.Marshal(httpResponse)
	if err != nil {
		return nil, err
	}

	return respJson, nil
}
