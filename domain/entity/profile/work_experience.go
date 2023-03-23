package profile

import "time"

type WorkExperience struct {
	id              string
	skillExperience string
	name            string
	dateRange       *DateRangeWork
	description     string
}
type DateRangeWork struct {
	start time.Time
	end   time.Time
}

type WorkExperienceDTO struct {
	Id              string
	SkillExperience string
	Name            string
	DateRange       *DateRangeWorkDTO
	Description     string
}
type DateRangeWorkDTO struct {
	Start time.Time
	End   time.Time
}

func NewWorkExperience(dto *WorkExperienceDTO) *WorkExperience {

	dateRange := NewWorkDateRange(dto.DateRange)
	return &WorkExperience{
		id:              dto.Id,
		skillExperience: dto.SkillExperience,
		name:            dto.Name,
		dateRange:       dateRange,
		description:     dto.Description,
	}
}

func NewWorkDateRange(dto *DateRangeWorkDTO) *DateRangeWork {
	return &DateRangeWork{ //Validation
		start: dto.Start,
		end:   dto.End,
	}
}
