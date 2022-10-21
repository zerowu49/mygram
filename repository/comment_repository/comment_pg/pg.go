package comment_pg

import (
	"mygram/entity"
	"mygram/pkg/errs"
	"mygram/repository/comment_repository"
	"time"

	"gorm.io/gorm"
)

type commentPG struct {
	db *gorm.DB
}

func NewCommentPG(db *gorm.DB) comment_repository.CommentRepository {
	return &commentPG{db: db}
}

func (c *commentPG) PostComment(commentPayload *entity.Comment) (*entity.Comment, errs.MessageErr) {
	comment := entity.Comment{}
	comment.UserID = commentPayload.UserID

	if err := c.db.Model(&comment).Create(&commentPayload).Error; err != nil {
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	if err := c.db.Last(&comment).Error; err != nil {
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	return &comment, nil
}

func (c *commentPG) GetAllComments() ([]*entity.Comment, errs.MessageErr) {
	comments := []*entity.Comment{}

	err := c.db.Preload("User").Preload("Photo").Find(&comments).Error
	if err != nil {
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	return comments, nil
}

func (c *commentPG) GetCommentByID(commentID uint) (*entity.Comment, errs.MessageErr) {
	comment := entity.Comment{}

	if err := c.db.Where("id = ?", commentID).First(&comment).Error; err != nil {
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	return &comment, nil
}

func (c *commentPG) EditCommentData(commentID uint, commentPayload *entity.Comment) (*entity.Comment, errs.MessageErr) {
	comment := entity.Comment{}

	err := c.db.Raw("Update comments SET message = ?, updated_at = ? WHERE id = ? RETURNING id, message, photo_id, user_id, created_at, updated_at", commentPayload.Message, time.Now(), commentID).Take(&comment).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.NewNotFoundError("Comment not found")
		}
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	return &comment, nil
}

func (c *commentPG) DeleteComment(commentID uint) errs.MessageErr {
	comment := entity.Comment{}

	if err := c.db.Where("id = ?", commentID).Delete(&comment).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errs.NewNotFoundError("Comment not found")
		}
		return errs.NewInternalServerErrorr("Something went wrong")
	}

	return nil
}
