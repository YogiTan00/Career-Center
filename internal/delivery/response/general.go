package response

import "encoding/json"

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type CustomReponseSingle struct {
	Status *Status `json:"status"`
}

type CustomReponseSingleInterface struct {
	Status *Status     `json:"status"`
	Data   interface{} `json:"data"`
}

type CustomReponseSingleInterfaceWithCount struct {
	Status *Status     `json:"status"`
	Data   interface{} `json:"data"`
	Count  int         `json:"count"`
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

func MapResponseInterface(code int, message string, data interface{}) ([]byte, error) {
	httpResponse := &CustomReponseSingleInterface{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: data,
	}

	respJson, err := json.Marshal(httpResponse)
	if err != nil {
		return nil, err
	}

	return respJson, nil
}

func MapResponseInterfaceWithCount(code int, message string, data interface{}, count int) ([]byte, error) {
	httpResponse := &CustomReponseSingleInterfaceWithCount{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data:  data,
		Count: count,
	}

	respJson, err := json.Marshal(httpResponse)
	if err != nil {
		return nil, err
	}

	return respJson, nil
}
