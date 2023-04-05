package profile

import "time"

type Education struct {
	id              string
	name            string
	dateRange       DateRangeEdu
	skillExperience string
	description     string
}
type DateRangeEdu struct {
	start time.Time
	end   time.Time
}

type EducationDTO struct {
	Id              string
	Name            string
	DateRange       DateRangeEduDTO
	SkillExperience string
	Description     string
}
type DateRangeEduDTO struct {
	Start time.Time
	End   time.Time
}

func NewEducation(dto *EducationDTO) *Education {
	dateRange := NewEduDateRange(dto.DateRange)
	return &Education{
		id:              dto.Id,
		name:            dto.Name,
		dateRange:       dateRange,
		skillExperience: dto.SkillExperience,
		description:     dto.Description,
	}
}

func NewEduDateRange(dto DateRangeEduDTO) DateRangeEdu {
	return DateRangeEdu{ //Validation
		start: dto.Start,
		end:   dto.End,
	}
}

func NewListEducation(dto []*EducationDTO) []*Education {
	listEducation := make([]*Education, 0)
	for _, data := range dto {
		education := NewEducation(data)
		listEducation = append(listEducation, education)
	}
	return listEducation
}
