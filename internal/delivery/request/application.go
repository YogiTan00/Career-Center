package request

import (
	"CareerCenter/domain/entity"
	"CareerCenter/domain/valueobject"
)

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

type RequestStatusApplicant struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}

func NewApplicantRequest(req *RequestStatusApplicant) *entity.StatusApplicantRequest {
	status := valueobject.NewTypeStatusApplicantFromStringPointer(req.Status)
	return &entity.StatusApplicantRequest{
		Id:     req.Id,
		Status: status,
	}
}
