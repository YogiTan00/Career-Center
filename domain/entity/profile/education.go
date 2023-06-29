package profile

import (
	"CareerCenter/domain/valueobject"
	"CareerCenter/utils/exceptions"
	"time"

	"github.com/google/uuid"
)

type Education struct {
	id             string
	email          string
	level          valueobject.TypeLevel
	name           string
	major          string
	stillEducation bool
	dateRange      DateRangeEdu
	description    string
}
type DateRangeEdu struct {
	start time.Time
	end   time.Time
}

type EducationDTO struct {
	Id             string
	Email          string
	Level          valueobject.TypeLevel
	Name           string
	Major          string
	StillEducation bool
	DateRange      DateRangeEduDTO
	Description    string
}
type DateRangeEduDTO struct {
	Start time.Time
	End   time.Time
}

func NewEducation(dto *EducationDTO) (*Education, error) {
	err := dto.Validation()
	if err != nil {
		return nil, err
	}
	dateRange := NewEduDateRange(dto.DateRange)
	level := valueobject.NewTypeLevelFromString(dto.Major)
	return &Education{
		id:             dto.Id,
		email:          dto.Email,
		level:          level,
		name:           dto.Name,
		major:          dto.Major,
		dateRange:      dateRange,
		stillEducation: dto.StillEducation,
		description:    dto.Description,
	}, nil
}

func NewEduDateRange(dto DateRangeEduDTO) DateRangeEdu {
	return DateRangeEdu{ //Validation
		start: dto.Start,
		end:   dto.End,
	}
}

func NewListEducation(dto []*EducationDTO) ([]*Education, error) {
	listEducation := make([]*Education, 0)
	for _, data := range dto {
		education, err := NewEducation(data)
		if err != nil {
			return nil, err
		}
		listEducation = append(listEducation, education)
	}
	return listEducation, nil
}

func (dto *EducationDTO) Validation() error {
	if len(dto.Id) <= 0 {
		dto.Id = uuid.NewString()
	}

	if dto.DateRange.Start.IsZero() {
		return exceptions.ErrorStartDate
	}

	if dto.DateRange.End.IsZero() && !dto.StillEducation {
		return exceptions.ErrCustomString("still education cant be false")
	}

	if !dto.DateRange.Start.IsZero() && !dto.DateRange.End.IsZero() {
		if dto.DateRange.End.Unix() < dto.DateRange.Start.Unix() {
			return exceptions.ErrorEndDate
		}
	}

	if dto.Level.StringLevel() == "" {
		return exceptions.ErrCustomString("error level education")
	}
	return nil
}

func (data *Education) GetId() string {
	return data.id
}
func (data *Education) GetEmail() string {
	return data.email
}
func (data *Education) GetLevel() valueobject.TypeLevel {
	return data.level
}
func (data *Education) GetName() string {
	return data.name
}
func (data *Education) GetMajor() string {
	return data.major
}
func (data *Education) GetStillEducation() bool {
	return data.stillEducation
}
func (data *Education) GetStartEducation() time.Time {
	return data.dateRange.start
}
func (data *Education) GetEndEducation() time.Time {
	return data.dateRange.end
}
func (data *Education) GetDescription() string {
	return data.description
}

func (data *Education) SetEmail(email string) {
	data.email = email
}
