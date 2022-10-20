package photo_repository

import (
	"mygram/entity"
	"mygram/pkg/errs"
)

type PhotoRepository interface {
	PostPhoto(photo *entity.Photo) (*entity.Photo, errs.MessageErr)
	GetAllPhotos() ([]*entity.Photo, errs.MessageErr)
	GetPhotoByID(photoID uint) (*entity.Photo, errs.MessageErr)
	EditPhotoData(photoID uint, photo *entity.Photo) (*entity.Photo, errs.MessageErr)
	DeletePhoto(photoID uint) errs.MessageErr
}
