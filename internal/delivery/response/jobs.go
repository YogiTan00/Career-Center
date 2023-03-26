package response

import (
	"CareerCenter/domain/entity"
	"time"
)

type JobsResponse struct {
	Id        string    `json:"id"`
	Position  string    `json:"position"`
	Company   string    `json:"company"`
	Logo      string    `json:"logo"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"createdAt"`
}

type DetailJobResponse struct {
	Id             string    `json:"id"`
	Position       string    `json:"position"`
	Company        string    `json:"company"`
	Logo           string    `json:"logo"`
	Address        string    `json:"address"`
	Status         bool      `json:"status"`
	SendDate       time.Time `json:"sendDate"`
	Qualification  string    `json:"qualification"`
	JobDescription string    `json:"jobDescription"`
	Category       string    `json:"category"`
	Description    string    `json:"description"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

func GetJobResponse(dto *entity.JobsDTO) *JobsResponse {
	return &JobsResponse{
		Id:        dto.Id,
		Position:  dto.Position,
		Company:   dto.Company,
		Logo:      dto.Logo,
		Address:   dto.Address,
		CreatedAt: dto.CreatedAt,
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
		Position:       dto.Position,
		Company:        dto.Company,
		Logo:           dto.Logo,
		Address:        dto.Address,
		Status:         dto.Status,
		SendDate:       dto.SendDate,
		Qualification:  dto.Qualification,
		JobDescription: dto.JobDescription,
		Category:       dto.Category,
		Description:    dto.Description,
		CreatedAt:      dto.CreatedAt,
		UpdatedAt:      dto.UpdatedAt,
	}
}
