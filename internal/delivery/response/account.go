package response

import (
	"encoding/json"
)

type Status struct {
	Code    int
	Message string
}

type CustomReponseSingle struct {
	Status *Status
}

type CustomReponseSingleAuth struct {
	Status *Status
	Data   string
}

func MapResponse(code int, message string) ([]byte, error) {
	httpResponse := &CustomReponseSingle{
		Status: &Status{
			Code:    code,
			Message: message,
		},
	}

	respJson, err := json.Marshal(httpResponse)
	if err != nil {
		return nil, err
	}

	return respJson, nil
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
