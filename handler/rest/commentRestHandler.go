package rest

import (
	"net/http"

	"mygram/dto"
	"mygram/entity"
	"mygram/pkg/helpers"
	"mygram/service"

	"github.com/gin-gonic/gin"
)

type commentRestHandler struct {
	commentService service.CommentService
}

func NewCommentRestHandler(commentService service.CommentService) *commentRestHandler {
	return &commentRestHandler{commentService: commentService}
}

// PostComment godoc
// @Schemes http
// @Tags Comment
// @Description Create new comment
// @Param        RequestBody   body      dto.CommentRequest  true  "User Request"
// @Param Authorization header string true "Authorization" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFAYS5jb20iLCJleHAiOjE2NjYzNjc3MjcsImlkIjoyfQ.T5NaUHAkGra7ehNfJ09rCLUdz9uDfjTYSbj1EtSbg5Y)
// @Accept json
// @Produce json
// @Success 201 {object} dto.CommentResponse
// @Failure 400 {object} object
// @Router /comments [post]
func (c *commentRestHandler) PostComment(ctx *gin.Context) {
	var commentRequest dto.CommentRequest
	var err error
	var userData entity.User

	contentType := helpers.GetContentType(ctx)
	if contentType == helpers.AppJSON {
		err = ctx.ShouldBindJSON(&commentRequest)
	} else {
		err = ctx.ShouldBind(&commentRequest)
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad_request",
			"message": err.Error(),
		})
		return
	}

	if value, ok := ctx.MustGet("userData").(entity.User); !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"err_message": "unauthorized",
		})
		return
	} else {
		userData = value
	}

	comment, err2 := c.commentService.PostComment(userData.ID, &commentRequest)
	if err2 != nil {
		ctx.JSON(err2.Status(), gin.H{
			"error":   err2.Error(),
			"message": err2.Message(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, comment)
}

// GetAllComments godoc
// @Schemes http
// @Tags Comment
// @Description Get All comment by identify user from token
// @Param Authorization header string true "Authorization" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFAYS5jb20iLCJleHAiOjE2NjYzNjc3MjcsImlkIjoyfQ.T5NaUHAkGra7ehNfJ09rCLUdz9uDfjTYSbj1EtSbg5Y)
// @Accept json
// @Produce json
// @Success 200 {array} []dto.GetCommentResponse
// @Failure 400 {object} object
// @Router /comments [get]
func (c *commentRestHandler) GetAllComments(ctx *gin.Context) {
	var userData entity.User
	if value, ok := ctx.MustGet("userData").(entity.User); !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"err_message": "unauthorized",
		})
		return
	} else {
		userData = value
	}
	_ = userData

	comments, err := c.commentService.GetAllComments()
	if err != nil {
		ctx.JSON(err.Status(), gin.H{
			"error":   err.Error(),
			"message": err.Message(),
		})
		return
	}

	ctx.JSON(http.StatusOK, comments)
}

// UpdateComment godoc
// @Schemes http
// @Tags Comment
// @Description Update comment by identify user from token
// @Param        RequestBody   body      dto.EditCommentRequest  true  "User Request"
// @Param commentID path string true "Comment ID" default(1)
// @Param Authorization header string true "Authorization" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFAYS5jb20iLCJleHAiOjE2NjYzNjc3MjcsImlkIjoyfQ.T5NaUHAkGra7ehNfJ09rCLUdz9uDfjTYSbj1EtSbg5Y)
// @Accept json
// @Produce json
// @Success 200 {object} dto.EditCommentResponse
// @Failure 400 {object} object
// @Router /comments/{commentID} [put]
func (c *commentRestHandler) UpdateComment(ctx *gin.Context) {
	var commentRequest dto.EditCommentRequest
	var userData entity.User
	var err error

	contentType := helpers.GetContentType(ctx)
	if contentType == helpers.AppJSON {
		err = ctx.ShouldBindJSON(&commentRequest)
	} else {
		err = ctx.ShouldBind(&commentRequest)
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad_request",
			"message": err.Error(),
		})
		return
	}

	if value, ok := ctx.MustGet("userData").(entity.User); !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"err_message": "unauthorized",
		})
		return
	} else {
		userData = value
	}
	_ = userData

	commentIdParam, err := helpers.GetParamId(ctx, "commentID")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err_message": "invalid params",
		})
		return
	}

	comment, err2 := c.commentService.EditCommentData(commentIdParam, &commentRequest)
	if err2 != nil {
		ctx.JSON(err2.Status(), gin.H{
			"error":   err2.Error(),
			"message": err2.Message(),
		})
		return
	}

	ctx.JSON(http.StatusOK, comment)
}

// DeleteComment godoc
// @Schemes http
// @Tags Comment
// @Description Delete comment by identify user from token
// @Param commentID path string true "Comment ID" default(2)
// @Param Authorization header string true "Authorization" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFAYS5jb20iLCJleHAiOjE2NjYzNjc3MjcsImlkIjoyfQ.T5NaUHAkGra7ehNfJ09rCLUdz9uDfjTYSbj1EtSbg5Y)
// @Accept json
// @Produce json
// @Success 200 {object} dto.DeleteCommentResponse
// @Failure 400 {object} object
// @Router /comments/{commentID} [delete]
func (c *commentRestHandler) DeleteComment(ctx *gin.Context) {
	var userData entity.User
	if value, ok := ctx.MustGet("userData").(entity.User); !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"err_message": "unauthorized",
		})
		return
	} else {
		userData = value
	}
	_ = userData

	commentIdParam, err := helpers.GetParamId(ctx, "commentID")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err_message": "invalid params",
		})
		return
	}

	res, err2 := c.commentService.DeleteComment(commentIdParam)
	if err2 != nil {
		ctx.JSON(err2.Status(), gin.H{
			"error":   err2.Error(),
			"message": err2.Message(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
