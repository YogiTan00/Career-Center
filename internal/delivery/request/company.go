package request

import (
	"CareerCenter/domain/entity"
	"CareerCenter/domain/valueobject"
)

type RequestCompany struct {
	EmailUser   string        `json:"emailUser"`
	Name        string        `json:"name"`
	TypeCompany string        `json:"typeCompany"`
	Address     string        `json:"address"`
	About       *RequestAbout `json:"about"`
}

type RequestAbout struct {
	Profile  string `json:"profile"`
	Website  string `json:"website"`
	Location string `json:"location"`
}

func NewCompanyRequest(req *RequestCompany) *entity.CompanyDTO {
	typeCompany := valueobject.NewTypeCompanyFromString(req.TypeCompany)
	return &entity.CompanyDTO{
		Email:       req.EmailUser,
		Name:        req.Name,
		TypeCompany: typeCompany,
		Address:     req.Address,
		About: &entity.AboutDTO{
			Profile:  req.About.Profile,
			Website:  req.About.Website,
			Location: req.About.Location,
		},
	}
}
