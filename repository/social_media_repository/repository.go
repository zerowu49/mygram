package social_media_repository

import (
	"mygram/entity"
	"mygram/pkg/errs"
)

type SocialMediaRepository interface {
	AddSocialMedia(socialMedia *entity.SocialMedia) (*entity.SocialMedia, errs.MessageErr)
	GetAllSocialMedias() ([]*entity.SocialMedia, errs.MessageErr)
	GetSocialMediaByID(socialMediaID uint) (*entity.SocialMedia, errs.MessageErr)
	EditSocialMediaData(socialMediaID uint, socialMedia *entity.SocialMedia) (*entity.SocialMedia, errs.MessageErr)
	DeleteSocialMedia(socialMediaID uint) errs.MessageErr
}
