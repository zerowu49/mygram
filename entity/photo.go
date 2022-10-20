package entity

import (
	"mygram/dto"
)

type Photo struct {
	GormModel
	Title    string `gorm:"not null" form:"title" json:"title" valid:"required~Title is required"`
	Caption  string `form:"caption" json:"caption" valid:"required~Caption is required"`
	PhotoURL string `gorm:"not null" form:"photo_url" json:"photo_url" valid:"required~Photo URL is required"`
	UserID   uint   `json:"user_id"`
	User     *User
}

func (p *Photo) ToGetPhotoResponseDTO() *dto.GetPhotoResponse {
	return &dto.GetPhotoResponse{
		ID:        p.ID,
		Title:     p.Title,
		Caption:   p.Caption,
		PhotoURL:  p.PhotoURL,
		UserID:    p.UserID,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		User: dto.EmbeddedUserResponse{
			Username: p.User.Username,
			Email:    p.User.Email,
		},
	}
}
