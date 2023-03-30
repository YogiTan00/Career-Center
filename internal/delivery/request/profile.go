package request

import (
	"CareerCenter/domain/entity/profile"
)

type RequestUpdateProfile struct {
	Nama        string `json:"name"`
	Skill       string `json:"skill"`
	PhoneNumber string `json:"phoneNumber"`
}

func NewUpdateProfileRequest(req *RequestUpdateProfile) *profile.ProfileUserDTO {
	return &profile.ProfileUserDTO{
		Name:        req.Nama,
		Skill:       req.Skill,
		PhoneNumber: req.PhoneNumber,
	}
}
