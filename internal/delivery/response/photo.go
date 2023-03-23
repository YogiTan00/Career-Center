package response

import (
	"CareerCenter/domain/entity/profile"
	"time"
)

type PhotoResponse struct {
	Id        string    `json:"id"`
	Filename  string    `json:"filename"`
	Size      int       `json:"size"`
	CreatedAt time.Time `json:"createdAt"`
}

func GetPhoto(dto *profile.PhotoProfileDTO) *PhotoResponse {
	return &PhotoResponse{
		Id:        dto.Id,
		Filename:  dto.Filename,
		Size:      dto.Size,
		CreatedAt: dto.CreatedAt,
	}
}
