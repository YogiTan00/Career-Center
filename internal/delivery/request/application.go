package request

import "CareerCenter/domain/entity"

type RequestApplication struct {
	CompanyId string `json:"companyId"`
	JobId     string `json:"jobId"`
}

func NewApplicationRequest(req *RequestApplication) *entity.ApplicationRequest {
	return &entity.ApplicationRequest{
		CompanyId: req.CompanyId,
		JobId:     req.JobId,
	}
}
