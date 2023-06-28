package testdata

import (
	"CareerCenter/domain/entity/profile"
	"CareerCenter/domain/valueobject"
	"CareerCenter/utils"
	"time"

	uuid2 "github.com/google/uuid"
)

func TestDataProfile(countWorkExperiencet int, countEducation int) *profile.ProfileUserDTO {
	var (
		listWorkExperience []*profile.WorkExperienceDTO
		listEducation      []*profile.EducationDTO
	)

	test := utils.RandomString(4)
	id, _ := uuid2.NewUUID()

	for i := 0; i < countWorkExperiencet; i++ {
		workExperience := &profile.WorkExperienceDTO{
			Id:              id.String(),
			Email:           test + "@gmail.com",
			SkillExperience: "DKV",
			Name:            "Toko Super 2013",
			StillWorking:    false,
			DateRange: profile.DateRangeWorkDTO{
				Start: time.Now().AddDate(-2, 0, 0),
				End:   time.Now().AddDate(-1, 0, 0),
			},
			Description: "pengalaman kerja di salt",
		}
		listWorkExperience = append(listWorkExperience, workExperience)
	}
	for i := 0; i < countEducation; i++ {
		education := &profile.EducationDTO{
			Id:             id.String(),
			Email:          test + "@gmail.com",
			Level:          valueobject.TypeLevel{},
			Name:           "SMA Lulusan 1023",
			Major:          "IT",
			StillEducation: false,
			DateRange: profile.DateRangeEduDTO{
				Start: time.Now().AddDate(-5, 0, 0),
				End:   time.Now().AddDate(-2, 0, 0),
			},
			Description: "sekolah mengenah atas",
		}
		listEducation = append(listEducation, education)
	}

	profile := &profile.ProfileUserDTO{
		Name:           test,
		Photo:          test + "@gmail.com.png",
		Skill:          "Programer Golang",
		Email:          test + "@gmail.com",
		PhoneNumber:    "08123123",
		WorkExperience: listWorkExperience,
		Education:      listEducation,
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
		UpdatedAt:  time.Now(),
	}
	return profile
}
