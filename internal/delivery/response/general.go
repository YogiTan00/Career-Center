package response

import "encoding/json"

type Status struct {
	Code    int
	Message string
}

type CustomReponseSingle struct {
	Status *Status
}

type CustomReponseSingleInterface struct {
	Status *Status
	Data   interface{}
}

type CustomReponseSingleInterfaceWithCount struct {
	Status *Status
	Data   interface{}
	Count  int
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
