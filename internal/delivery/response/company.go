package response

import (
	"CareerCenter/domain/entity"
	"time"
)

type CompanyResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	TypeCompany string `json:"typeCompany"`
	Address     string `json:"address"`
	Logo        string `json:"logo"`
}

type CompanyProfileResponse struct {
	Id          string            `json:"id"`
	Name        string            `json:"name"`
	TypeCompany string            `json:"typeCompany"`
	Address     string            `json:"address"`
	Logo        string            `json:"logo"`
	About       *entity.AboutDTO  `json:"about"`
	Jobs        []*entity.JobsDTO `json:"jobs"`
	CreatedAt   time.Time         `json:"createdAt"`
	UpdatedAt   time.Time         `json:"updatedAt"`
}

func GetCompanyResponse(dto *entity.CompanyDTO) *CompanyResponse {
	return &CompanyResponse{
		Id:          dto.Id,
		Name:        dto.Name,
		TypeCompany: dto.TypeCompany.StringCompany(),
		Address:     dto.Address,
		Logo:        dto.Logo,
	}
}

func GetListCompanyResponse(dto []*entity.CompanyDTO) []*CompanyResponse {
	listCompany := make([]*CompanyResponse, 0)
	for _, data := range dto {
		job := GetCompanyResponse(data)
		listCompany = append(listCompany, job)
	}
	return listCompany
}

func GetCompanyProfileResponse(dto *entity.CompanyDTO, dtoJobs []*entity.JobsDTO) *CompanyProfileResponse {
	jobs := make([]*entity.JobsDTO, 0)
	for _, data := range dtoJobs {
		job := &entity.JobsDTO{
			Id:             data.Id,
			Position:       data.Position,
			Company:        data.Company,
			Logo:           data.Logo,
			Address:        data.Address,
			Status:         data.Status,
			SendDate:       data.SendDate,
			Qualification:  data.Qualification,
			JobDescription: data.JobDescription,
			Category:       data.Category,
			Description:    data.Description,
			CreatedAt:      data.CreatedAt,
			UpdatedAt:      data.UpdatedAt,
		}
		jobs = append(jobs, job)
	}
	return &CompanyProfileResponse{
		Id:          dto.Id,
		Name:        dto.Name,
		TypeCompany: dto.TypeCompany.StringCompany(),
		Address:     dto.Address,
		Logo:        dto.Logo,
		About: &entity.AboutDTO{
			Profile:  dto.About.Profile,
			Website:  dto.About.Website,
			Location: dto.About.Location,
		},
		Jobs:      jobs,
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
	}
}
