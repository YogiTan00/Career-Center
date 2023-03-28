package testdata

import (
	"CareerCenter/domain/entity/profile"
	"CareerCenter/utils"
	uuid2 "github.com/google/uuid"
	"time"
)

func TestDataProfile() *profile.ProfileUserDTO {
	test := utils.RandomString(4)
	id, _ := uuid2.NewUUID()
	profile := &profile.ProfileUserDTO{
		Name:        test,
		Photo:       test + "@gmail.com.png",
		Skill:       "Programer Golang",
		Email:       test + "@gmail.com",
		PhoneNumber: "08123123",
		WorkExperience: &profile.WorkExperienceDTO{
			Id:              id.String(),
			SkillExperience: "DKV",
			Name:            "Toko Super 2013",
			DateRange: &profile.DateRangeWorkDTO{
				Start: time.Now().AddDate(-2, 0, 0),
				End:   time.Now().AddDate(-1, 0, 0),
			},
			Description: "pengalaman kerja di salt",
		},
		Education: &profile.EducationDTO{
			Id:   id.String(),
			Name: "SMA Lulusan 1023",
			DateRange: &profile.DateRangeEduDTO{
				Start: time.Now().AddDate(-5, 0, 0),
				End:   time.Now().AddDate(-2, 0, 0),
			},
			SkillExperience: "Jaringan",
			Description:     "sekolah mengenah atas",
		},
		Ability: []string{
			"Microsoft Excel",
			"Microsoft Word",
			"Microsoft Power POint",
		},
		Language: []string{
			"English",
			"Indonesia",
			"Japan",
		},
		CvResume:   "CV " + test + "@gmail.com.pdf",
		Portofolio: "Portofolio " + test + "@gmail.com.pdf",
		CreatedAt:  time.Now(),
		UpdateAt:   time.Now(),
	}
	return profile
}
