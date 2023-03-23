package response

import (
	"CareerCenter/domain/entity"
	"CareerCenter/domain/entity/profile"
)

type HomeResponse struct {
	Name  string          `json:"name"`
	Skill string          `json:"skill"`
	Photo *PhotoResponse  `json:"photo"`
	Jobs  []*JobsResponse `json:"jobs"`
}

func GetHome(photo *profile.PhotoProfileDTO, jobs []*entity.JobsDTO, res *entity.HomePageDTO) *HomeResponse {
	photoProfile := GetPhoto(photo)
	listJobs := GetListJob(jobs)
	return &HomeResponse{
		Name:  res.Profile.Name,
		Skill: res.Profile.Skill,
		Photo: photoProfile,
		Jobs:  listJobs,
	}
}
