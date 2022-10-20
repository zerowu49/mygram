package rest

import (
	"net/http"

	"mygram/dto"
	"mygram/entity"
	"mygram/pkg/helpers"
	"mygram/service"

	"github.com/gin-gonic/gin"
)

type photoRestHandler struct {
	photoService service.PhotoService
}

func NewPhotoRestHandler(photoService service.PhotoService) *photoRestHandler {
	return &photoRestHandler{photoService: photoService}
}

// PostPhoto godoc
// @Schemes http
// @Tags Photos
// @Description Post Photo based on token user
// @Param        RequestBody   body      dto.PhotoRequest  true  "User Request"
// @Accept json
// @Produce json
// @Success 201 {object} dto.PhotoResponse
// @Failure 422 {object} object
// @param Authorization header string true "Authorization" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFAYS5jb20iLCJleHAiOjE2NjYyODU5MzYsImlkIjoyfQ.ahlQ3Icb2h72Aam9jWgZFM-TultwjKzymM7DGKHVVKU)
// @Router /photos [post]
func (p *photoRestHandler) PostPhoto(c *gin.Context) {
	var photoRequest dto.PhotoRequest
	var err error
	var userData entity.User

	contentType := helpers.GetContentType(c)
	if contentType == helpers.AppJSON {
		err = c.ShouldBindJSON(&photoRequest)
	} else {
		err = c.ShouldBind(&photoRequest)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad_request",
			"message": err.Error(),
		})
		return
	}

	if value, ok := c.MustGet("userData").(entity.User); !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err_message": "unauthorized",
		})
		return
	} else {
		userData = value
	}

	photo, err := p.photoService.PostPhoto(userData.ID, &photoRequest)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   "unprocessable_entity",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, photo)
}

// GetAllPhotos godoc
// @Schemes http
// @Tags Photos
// @Description Get All Photos based on token user
// @Accept json
// @Produce json
// @Success 200 {array} dto.GetPhotoResponse
// @Failure 403 {object} object
// @param Authorization header string true "Authorization" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFAYS5jb20iLCJleHAiOjE2NjYyODU5MzYsImlkIjoyfQ.ahlQ3Icb2h72Aam9jWgZFM-TultwjKzymM7DGKHVVKU)
// @Router /photos [get]
func (p *photoRestHandler) GetAllPhotos(c *gin.Context) {
	var userData entity.User
	if value, ok := c.MustGet("userData").(entity.User); !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err_message": "unauthorized",
		})
		return
	} else {
		userData = value
	}
	_ = userData

	photos, err := p.photoService.GetAllPhotos()
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error":   "forbidden",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, photos)
}

// UpdatePhoto godoc
// @Schemes http
// @Tags Photos
// @Description Update Photo based on token user with paramId of photoId
// @Param        RequestBody   body      dto.PhotoRequest  true  "User Request"
// @Accept json
// @Produce json
// @Success 202 {object} dto.PhotoUpdateResponse
// @Failure 403 {object} object
// @Param Authorization header string true "Authorization" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFAYS5jb20iLCJleHAiOjE2NjYyODU5MzYsImlkIjoyfQ.ahlQ3Icb2h72Aam9jWgZFM-TultwjKzymM7DGKHVVKU)
// @Param photoID path string true "Photo ID" default(1)
// @Router /photos/{photoID} [put]
func (p *photoRestHandler) UpdatePhoto(c *gin.Context) {
	var photoRequest dto.PhotoRequest
	var err error

	contentType := helpers.GetContentType(c)
	if contentType == helpers.AppJSON {
		err = c.ShouldBindJSON(&photoRequest)
	} else {
		err = c.ShouldBind(&photoRequest)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad_request",
			"message": err.Error(),
		})
		return
	}

	photoIdParam, err := helpers.GetParamId(c, "photoID")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err_message": "invalid params",
		})
		return
	}

	photo, err2 := p.photoService.EditPhotoData(photoIdParam, &photoRequest)
	if err2 != nil {
		c.JSON(err2.Status(), gin.H{
			"error":   err2.Error(),
			"message": err2.Message(),
		})
		return
	}

	c.JSON(http.StatusAccepted, photo)
}

// DeletePhoto godoc
// @Schemes http
// @Tags Photos
// @Description Delete Photo based on token user with paramId of photoId
// @Param photoID path string true "Photo ID" default(2)
// @Accept json
// @Produce json
// @Success 200 {object} dto.DeletePhotoResponse
// @Failure 403 {object} object
// @param Authorization header string true "Authorization" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFAYS5jb20iLCJleHAiOjE2NjYyODU5MzYsImlkIjoyfQ.ahlQ3Icb2h72Aam9jWgZFM-TultwjKzymM7DGKHVVKU)
// @Router /photos/{photoID} [delete]
func (p *photoRestHandler) DeletePhoto(c *gin.Context) {
	photoIdParam, err := helpers.GetParamId(c, "photoID")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err_message": "invalid params",
		})
		return
	}

	_ = photoIdParam

	res, err2 := p.photoService.DeletePhoto(photoIdParam)
	if err2 != nil {
		c.JSON(err2.Status(), gin.H{
			"error":   err2.Error(),
			"message": err2.Message(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}
