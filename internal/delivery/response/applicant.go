package response

import "CareerCenter/domain/entity"

type ResponseApplication struct {
	Id          string `json:"id"`
	CompanyId   string `json:"companyId"`
	JobId       string `json:"jobId"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	Skill       string `json:"skill"`
	PhoneNumber string `json:"phoneNumber"`
	CvResume    string `json:"cvResume"`
	Portofolio  string `json:"portofolio"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	DeletedAt   string `json:"deletedAt"`
	Status      string `json:"status"`
}

func ResponseApplicantByJobId(dto *entity.ApplicationDTO) *ResponseApplication {
	return &ResponseApplication{
		Id:          dto.Id,
		CompanyId:   dto.CompanyId,
		JobId:       dto.JobId,
		Email:       dto.Email,
		Name:        dto.Name,
		Skill:       dto.Skill,
		PhoneNumber: dto.PhoneNumber,
		CvResume:    dto.CvResume,
		Portofolio:  dto.Portofolio,
		CreatedAt:   dto.CreatedAt.String(),
		UpdatedAt:   dto.UpdatedAt.String(),
		DeletedAt:   dto.DeletedAt.String(),
		Status:      dto.Status.String(),
	}
}

func ResponseListApplicantByJobId(dto []*entity.ApplicationDTO) []*ResponseApplication {
	res := make([]*ResponseApplication, 0)
	for _, data := range dto {
		applicant := ResponseApplicantByJobId(data)
		res = append(res, applicant)
	}
	return res
}
