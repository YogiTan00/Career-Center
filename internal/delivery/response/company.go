package response

import (
	"CareerCenter/domain/entity"
)

type CompanyResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	TypeCompany string `json:"typeCompany"`
	Address     string `json:"address"`
	Logo        string `json:"logo"`
}

type CompanyProfileResponse struct {
	Id          string               `json:"id"`
	Name        string               `json:"name"`
	TypeCompany string               `json:"typeCompany"`
	Address     string               `json:"address"`
	Logo        string               `json:"logo"`
	About       *AboutResponse       `json:"about"`
	Jobs        []*DetailJobResponse `json:"jobs"`
	CreatedAt   string               `json:"createdAt"`
	UpdatedAt   string               `json:"updatedAt"`
}

type AboutResponse struct {
	Profile  string `json:"profile"`
	Website  string `json:"website"`
	Location string `json:"location"`
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
	jobs := make([]*DetailJobResponse, 0)
	for _, data := range dtoJobs {
		job := &DetailJobResponse{
			Id:             data.Id,
			Position:       data.Position,
			Company:        data.Company,
			Logo:           data.Logo,
			Address:        data.Address,
			Status:         data.Status,
			Qualification:  data.Qualification,
			JobDescription: data.JobDescription,
			CreatedAt:      data.CreatedAt.String(),
			UpdatedAt:      data.UpdatedAt.String(),
		}
		jobs = append(jobs, job)
	}
	return &CompanyProfileResponse{
		Id:          dto.Id,
		Name:        dto.Name,
		TypeCompany: dto.TypeCompany.StringCompany(),
		Address:     dto.Address,
		Logo:        dto.Logo,
		About: &AboutResponse{
			Profile:  dto.About.Profile,
			Website:  dto.About.Website,
			Location: dto.About.Location,
		},
		Jobs:      jobs,
		CreatedAt: dto.CreatedAt.String(),
		UpdatedAt: dto.UpdatedAt.String(),
	}
}
