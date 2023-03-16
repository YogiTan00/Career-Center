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

func GetJob(dto *entity.JobsDTO) *JobsResponse {
	return &JobsResponse{
		Id:        dto.Id,
		Position:  dto.Position,
		Company:   dto.Company,
		Logo:      dto.Logo,
		Address:   dto.Address,
		CreatedAt: dto.CreatedAt,
	}
}

func GetListJob(dto []*entity.JobsDTO) []*JobsResponse {
	listJobs := make([]*JobsResponse, 0)
	for _, data := range dto {
		job := GetJob(data)
		listJobs = append(listJobs, job)
	}
	return listJobs
}
