package rest

import (
	"net/http"

	"mygram/dto"
	"mygram/entity"
	"mygram/pkg/helpers"
	"mygram/service"

	"github.com/gin-gonic/gin"
)

type socialMediaRestHandler struct {
	socialMediaService service.SocialMediaService
}

func NewSocialMediaRestHandler(socialMediaService service.SocialMediaService) *socialMediaRestHandler {
	return &socialMediaRestHandler{socialMediaService: socialMediaService}
}

// AddSocialMedia godoc
// @Schemes http
// @Tags Social Media
// @Description Add social media
// @Param        RequestBody   body      dto.SocialMediaRequest  true  "User Request"
// @Param Authorization header string true "Authorization" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFAYS5jb20iLCJleHAiOjE2NjYzNjc3MjcsImlkIjoyfQ.T5NaUHAkGra7ehNfJ09rCLUdz9uDfjTYSbj1EtSbg5Y)
// @Accept json
// @Produce json
// @Success 201 {object} dto.SocialMediaResponse
// @Failure 404 {object} object
// @Router /social-medias [post]
func (s *socialMediaRestHandler) AddSocialMedia(c *gin.Context) {
	var socialMediaRequest dto.SocialMediaRequest
	var err error
	var userData entity.User

	contentType := helpers.GetContentType(c)
	if contentType == helpers.AppJSON {
		err = c.ShouldBindJSON(&socialMediaRequest)
	} else {
		err = c.ShouldBind(&socialMediaRequest)
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

	socialMedia, err2 := s.socialMediaService.AddSocialMedia(userData.ID, &socialMediaRequest)
	if err2 != nil {
		c.JSON(err2.Status(), gin.H{
			"error":   err2.Error(),
			"message": err2.Message(),
		})
		return
	}

	c.JSON(http.StatusCreated, socialMedia)
}

// GetAllSocialMedias godoc
// @Schemes http
// @Tags Social Media
// @Description Add social media
// @Param Authorization header string true "Authorization" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFAYS5jb20iLCJleHAiOjE2NjYzNjc3MjcsImlkIjoyfQ.T5NaUHAkGra7ehNfJ09rCLUdz9uDfjTYSbj1EtSbg5Y)
// @Accept json
// @Produce json
// @Success 200 {array} dto.GetSocialMediaResponse
// @Failure 404 {object} object
// @Router /social-medias [get]
func (s *socialMediaRestHandler) GetAllSocialMedias(c *gin.Context) {
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

	socialMedia, err := s.socialMediaService.GetAllSocialMedias()
	if err != nil {
		c.JSON(err.Status(), gin.H{
			"error":   err.Error(),
			"message": err.Message(),
		})
		return
	}

	c.JSON(http.StatusOK, socialMedia)
}

// EditSocialMediaData godoc
// @Schemes http
// @Tags Social Media
// @Description Edit social media data
// @Param        RequestBody   body      dto.SocialMediaRequest  true  "User Request"
// @Param socmedId path string true "Social Media ID" default(1)
// @Param Authorization header string true "Authorization" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFAYS5jb20iLCJleHAiOjE2NjYzNjc3MjcsImlkIjoyfQ.T5NaUHAkGra7ehNfJ09rCLUdz9uDfjTYSbj1EtSbg5Y)
// @Accept json
// @Produce json
// @Success 200 {object} dto.SocialMediaResponse
// @Failure 404 {object} object
// @Router /social-medias/{socmedId} [put]
func (s *socialMediaRestHandler) EditSocialMediaData(c *gin.Context) {
	var socialMediaRequest dto.SocialMediaRequest
	var err error
	var userData entity.User

	contentType := helpers.GetContentType(c)
	if contentType == helpers.AppJSON {
		err = c.ShouldBindJSON(&socialMediaRequest)
	} else {
		err = c.ShouldBind(&socialMediaRequest)
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
	_ = userData

	socialMediaIdParam, err := helpers.GetParamId(c, "socialMediaID")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err_message": "invalid params",
		})
		return
	}

	socialMedia, err2 := s.socialMediaService.EditSocialMediaData(socialMediaIdParam, &socialMediaRequest)
	if err2 != nil {
		c.JSON(err2.Status(), gin.H{
			"error":   err2.Error(),
			"message": err2.Message(),
		})
		return
	}

	c.JSON(http.StatusOK, socialMedia)
}

// DeleteSocialMedia godoc
// @Schemes http
// @Tags Social Media
// @Description Delete social media data
// @Param socmedId path string true "Social Media ID" default(2)
// @Param Authorization header string true "Authorization" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFAYS5jb20iLCJleHAiOjE2NjYzNjc3MjcsImlkIjoyfQ.T5NaUHAkGra7ehNfJ09rCLUdz9uDfjTYSbj1EtSbg5Y)
// @Accept json
// @Produce json
// @Success 200 {object} dto.DeleteSocialMediaResponse
// @Failure 404 {object} object
// @Router /social-medias/{socmedId} [delete]
func (s *socialMediaRestHandler) DeleteSocialMedia(c *gin.Context) {
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

	socialMediaIdParam, err := helpers.GetParamId(c, "socialMediaID")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err_message": "invalid params",
		})
		return
	}

	res, err2 := s.socialMediaService.DeleteSocialMedia(socialMediaIdParam)
	if err2 != nil {
		c.JSON(err2.Status(), gin.H{
			"error":   err2.Error(),
			"message": err2.Message(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}
