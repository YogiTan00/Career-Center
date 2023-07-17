package response

import (
	"CareerCenter/domain/entity"
	"CareerCenter/utils"
)

type JobsResponse struct {
	Id        string `json:"id"`
	CompanyId string `json:"companyId"`
	Position  string `json:"position"`
	Company   string `json:"company"`
	Logo      string `json:"logo"`
	Address   string `json:"address"`
	Status    bool   `json:"status"`
	CreatedAt string `json:"createdAt"`
	Applicant int    `json:"applicant"`
}

type DetailJobResponse struct {
	Id             string `json:"id"`
	CompanyId      string `json:"companyId"`
	Position       string `json:"position"`
	Company        string `json:"company"`
	Logo           string `json:"logo"`
	Address        string `json:"address"`
	Status         bool   `json:"status"`
	ApplyDate      string `json:"applyDate"`
	Qualification  string `json:"qualification"`
	JobDescription string `json:"jobDescription"`
	CreatedAt      string `json:"createdAt"`
	UpdatedAt      string `json:"updatedAt"`
	Applicant      int    `json:"applicant"`
}

func GetJobResponse(dto *entity.JobsDTO) *JobsResponse {
	return &JobsResponse{
		Id:        dto.Id,
		CompanyId: dto.CompanyId,
		Position:  dto.Position,
		Company:   dto.Company,
		Logo:      dto.Logo,
		Address:   dto.Address,
		Status:    dto.Status,
		CreatedAt: utils.ToOnlyDateResponse(dto.CreatedAt),
		Applicant: dto.Applicant,
	}
}

func GetListJobResponse(dto []*entity.JobsDTO) []*JobsResponse {
	listJobs := make([]*JobsResponse, 0)
	for _, data := range dto {
		job := GetJobResponse(data)
		listJobs = append(listJobs, job)
	}
	return listJobs
}

func GetDetailJobResponse(dto *entity.JobsDTO) *DetailJobResponse {
	return &DetailJobResponse{
		Id:             dto.Id,
		CompanyId:      dto.CompanyId,
		Position:       dto.Position,
		Company:        dto.Company,
		Logo:           dto.Logo,
		Address:        dto.Address,
		Status:         dto.Status,
		ApplyDate:      utils.ToOnlyDateResponse(dto.ApplyDate),
		Qualification:  dto.Qualification,
		JobDescription: dto.JobDescription,
		CreatedAt:      utils.ToOnlyDateResponse(dto.CreatedAt),
		UpdatedAt:      utils.ToOnlyDateResponse(dto.UpdatedAt),
		Applicant:      dto.Applicant,
	}
}
