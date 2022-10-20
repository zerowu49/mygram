package service

import (
	"mygram/dto"
	"mygram/entity"
	"mygram/pkg/errs"
	"mygram/pkg/helpers"
	"mygram/repository/social_media_repository"
)

type SocialMediaService interface {
	AddSocialMedia(userID uint, socialMediaPayload *dto.SocialMediaRequest) (*dto.SocialMediaResponse, errs.MessageErr)
	GetAllSocialMedias() ([]*dto.GetSocialMediaResponse, errs.MessageErr)
	EditSocialMediaData(socialMediaID uint, socialMediaPayload *dto.SocialMediaRequest) (*dto.UpdateSocialMediaResponse, errs.MessageErr)
	DeleteSocialMedia(socialMediaID uint) (*dto.DeleteSocialMediaResponse, errs.MessageErr)
}

type socialMediaService struct {
	socialMediaRepository social_media_repository.SocialMediaRepository
}

func NewSocialMediaService(socialMediaRepository social_media_repository.SocialMediaRepository) SocialMediaService {
	return &socialMediaService{socialMediaRepository: socialMediaRepository}
}

func (s *socialMediaService) AddSocialMedia(userID uint, socialMediaPayload *dto.SocialMediaRequest) (*dto.SocialMediaResponse, errs.MessageErr) {
	if err := helpers.ValidateStruct(socialMediaPayload); err != nil {
		return nil, err
	}

	entity := entity.SocialMedia{
		Name:           socialMediaPayload.Name,
		SocialMediaURL: socialMediaPayload.SocialMediaURL,
		UserID:         userID,
	}

	socialMedia, err := s.socialMediaRepository.AddSocialMedia(&entity)
	if err != nil {
		return nil, err
	}

	response := &dto.SocialMediaResponse{
		ID:             socialMedia.ID,
		Name:           socialMedia.Name,
		SocialMediaURL: socialMedia.SocialMediaURL,
		UserID:         socialMedia.UserID,
		CreatedAt:      socialMedia.CreatedAt,
	}

	return response, nil
}

func (s *socialMediaService) GetAllSocialMedias() ([]*dto.GetSocialMediaResponse, errs.MessageErr) {
	dto := make([]*dto.GetSocialMediaResponse, 0)
	socialmedias, err := s.socialMediaRepository.GetAllSocialMedias()
	if err != nil {
		return nil, err
	}

	for _, socialmedia := range socialmedias {
		dto = append(dto, socialmedia.ToGetSocialMediaResponseDTO())
	}

	return dto, nil
}

func (s *socialMediaService) EditSocialMediaData(socialMediaID uint, socialMediaPayload *dto.SocialMediaRequest) (*dto.UpdateSocialMediaResponse, errs.MessageErr) {
	if err := helpers.ValidateStruct(socialMediaPayload); err != nil {
		return nil, err
	}

	entity := entity.SocialMedia{
		Name:           socialMediaPayload.Name,
		SocialMediaURL: socialMediaPayload.SocialMediaURL,
	}

	socialMedia, err := s.socialMediaRepository.EditSocialMediaData(socialMediaID, &entity)
	if err != nil {
		return nil, err
	}

	response := &dto.UpdateSocialMediaResponse{
		ID:             socialMedia.ID,
		Name:           socialMedia.Name,
		SocialMediaURL: socialMedia.SocialMediaURL,
		UserID:         socialMedia.UserID,
		UpdatedAt:      socialMedia.UpdatedAt,
	}

	return response, nil
}

func (s *socialMediaService) DeleteSocialMedia(socialMediaID uint) (*dto.DeleteSocialMediaResponse, errs.MessageErr) {
	if err := s.socialMediaRepository.DeleteSocialMedia(socialMediaID); err != nil {
		return nil, err
	}

	response := &dto.DeleteSocialMediaResponse{
		Message: "Your social media has been deleted",
	}

	return response, nil
}
