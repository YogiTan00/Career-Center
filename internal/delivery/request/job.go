package request

import (
	"CareerCenter/domain/entity"
)

type RequestJob struct {
	CompanyId      string `json:"companyId"`
	Position       string `json:"position"`
	Company        string `json:"company"`
	Logo           string `json:"logo"`
	Address        string `json:"address"`
	Status         bool   `json:"status"`
	Qualification  string `json:"qualification"`
	JobDescription string `json:"jobDescription"`
}

func NewJobRequest(req *RequestJob) *entity.JobsDTO {
	return &entity.JobsDTO{
		CompanyId:      req.CompanyId,
		Position:       req.Position,
		Company:        req.Company,
		Logo:           req.Logo,
		Address:        req.Address,
		Status:         req.Status,
		Qualification:  req.Qualification,
		JobDescription: req.JobDescription,
	}
}
