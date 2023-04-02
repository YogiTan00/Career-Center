package profile

import (
	"CareerCenter/domain/entity/account"
	"time"
)

type WorkExperience struct {
	id              string
	email           string
	skillExperience string
	name            string
	stillWorking    bool
	dateRange       DateRangeWork
	description     string
}
type DateRangeWork struct {
	start time.Time
	end   time.Time
}

type WorkExperienceDTO struct {
	Id              string
	Email           string
	SkillExperience string
	Name            string
	StillWorking    bool
	DateRange       DateRangeWorkDTO
	Description     string
}
type DateRangeWorkDTO struct {
	Start time.Time
	End   time.Time
}

func NewWorkExperience(dto WorkExperienceDTO) *WorkExperience {
	dateRange := NewWorkDateRange(dto.DateRange)
	return &WorkExperience{
		id:              dto.Id,
		email:           dto.Email,
		skillExperience: dto.SkillExperience,
		name:            dto.Name,
		stillWorking:    dto.StillWorking,
		dateRange:       dateRange,
		description:     dto.Description,
	}
}

func NewWorkExperienceByRegis(dto *account.Account) *WorkExperience {
	data := WorkExperienceDTO{
		Id:    dto.GetId(),
		Email: dto.GetEmail(),
	}
	result := NewWorkExperience(data)
	return result
}

func NewWorkDateRange(dto DateRangeWorkDTO) DateRangeWork {
	return DateRangeWork{ //Validation
		start: dto.Start,
		end:   dto.End,
	}
}

func (data *WorkExperience) GetId() string {
	return data.id
}
func (data *WorkExperience) GetEmail() string {
	return data.email
}
func (data *WorkExperience) GetSkillExperience() string {
	return data.skillExperience
}
func (data *WorkExperience) GetName() string {
	return data.name
}
func (data *WorkExperience) GetStillWorking() bool {
	return data.stillWorking
}
func (data *WorkExperience) GetStartWork() time.Time {
	return data.dateRange.start
}
func (data *WorkExperience) GetEndWork() time.Time {
	return data.dateRange.end
}
func (data *WorkExperience) GetDescription() string {
	return data.description
}
