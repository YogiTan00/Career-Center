package profile

import "time"

type PhotoProfile struct {
	id        string
	filename  string
	size      int
	createdAt time.Time
}

type PhotoProfileDTO struct {
	Id        string
	Filename  string
	Size      int
	CreatedAt time.Time
}

func NewPhoto(dto *PhotoProfileDTO) *PhotoProfile {
	return &PhotoProfile{ //Validation
		id:        dto.Id,
		filename:  dto.Filename,
		size:      dto.Size,
		createdAt: dto.CreatedAt,
	}
}
