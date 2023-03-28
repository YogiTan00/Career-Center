package request

type RequestApplication struct {
	CompanyId string `json:"companyId"`
}

func NewApplicationRequest(req *RequestApplication) string {
	return req.CompanyId
}
