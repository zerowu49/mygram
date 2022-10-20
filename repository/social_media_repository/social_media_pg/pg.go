package social_media_pg

import (
	"mygram/entity"
	"mygram/pkg/errs"
	"mygram/repository/social_media_repository"

	"gorm.io/gorm"
)

type socialMediaPG struct {
	db *gorm.DB
}

func NewSocialMediaPG(db *gorm.DB) social_media_repository.SocialMediaRepository {
	return &socialMediaPG{db: db}
}

func (s *socialMediaPG) AddSocialMedia(socialMediaPayload *entity.SocialMedia) (*entity.SocialMedia, errs.MessageErr) {
	socialMedia := entity.SocialMedia{}

	if err := s.db.Model(socialMedia).Create(&socialMediaPayload).Error; err != nil {
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	if err := s.db.Last(&socialMedia).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.NewNotFoundError("Social media not found")
		}
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	return &socialMedia, nil
}

func (s *socialMediaPG) GetAllSocialMedias() ([]*entity.SocialMedia, errs.MessageErr) {
	socialMedias := []*entity.SocialMedia{}

	if err := s.db.Preload("User").Find(&socialMedias).Error; err != nil {
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	return socialMedias, nil
}

func (s *socialMediaPG) GetSocialMediaByID(socialMediaID uint) (*entity.SocialMedia, errs.MessageErr) {
	socialMedia := entity.SocialMedia{}

	if err := s.db.Model(socialMedia).Where("id = ?", socialMediaID).First(&socialMedia).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.NewNotFoundError("Social media not found")
		}
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	return &socialMedia, nil
}

func (s *socialMediaPG) EditSocialMediaData(socialMediaID uint, socialMediaPayload *entity.SocialMedia) (*entity.SocialMedia, errs.MessageErr) {
	socialMedia := entity.SocialMedia{}

	if err := s.db.Model(socialMedia).Where("id = ?", socialMediaID).Updates(&socialMediaPayload).Take(&socialMedia).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.NewNotFoundError("Social media not found")
		}
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	return &socialMedia, nil
}

func (s *socialMediaPG) DeleteSocialMedia(socialMediaID uint) errs.MessageErr {
	socialMedia := entity.SocialMedia{}

	if err := s.db.Model(socialMedia).Where("id = ?", socialMediaID).Delete(&socialMedia).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errs.NewNotFoundError("Social media not found")
		}
		return errs.NewInternalServerErrorr("Something went wrong")
	}

	return nil
}
