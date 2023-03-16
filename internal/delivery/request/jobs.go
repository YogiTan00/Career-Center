package request

type RequestFilter struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}
