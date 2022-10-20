package rest

import (
	"fmt"
	"net/http"

	"mygram/dto"
	"mygram/entity"
	"mygram/pkg/helpers"
	"mygram/service"

	"github.com/gin-gonic/gin"
)

type userRestHandler struct {
	userService service.UserService
}

func NewUserRestHandler(userService service.UserService) *userRestHandler {
	return &userRestHandler{userService: userService}
}

// LoginUser godoc
// @Schemes http
// @Tags Users
// @Summary Login user
// @Description Login for existing user
// @Param        RequestBody   body      dto.LoginRequest  true  "User Request"
// @Accept json
// @Produce json
// @Success 201 {object} dto.LoginResponse
// @Failure 400 {object} object
// @Router /users/login [post]
func (u *userRestHandler) Login(c *gin.Context) {
	var userRequest dto.LoginRequest
	var err error

	contentType := helpers.GetContentType(c)
	if contentType == helpers.AppJSON {
		err = c.ShouldBindJSON(&userRequest)
	} else {
		err = c.ShouldBind(&userRequest)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad_request",
			"message": err.Error(),
		})
		return
	}

	token, err2 := u.userService.Login(&userRequest)

	if err2 != nil {
		c.JSON(err2.Status(), gin.H{
			"error":   err2.Error(),
			"message": err2.Message(),
		})
		return
	}
	c.JSON(http.StatusCreated, token)
}

// RegisterUser godoc
// @Schemes http
// @Tags Users
// @Summary Register new user
// @Description Register new user
// @Param        RequestBody   body      dto.RegisterRequest  true  "User Request"
// @Accept json
// @Produce json
// @Success 201 {object} dto.RegisterResponse
// @Failure 400 {object} object
// @Router /users/register [post]
func (u *userRestHandler) Register(c *gin.Context) {
	var userRequest dto.RegisterRequest
	var err error

	contentType := helpers.GetContentType(c)
	if contentType == helpers.AppJSON {
		err = c.ShouldBindJSON(&userRequest)
	} else {
		err = c.ShouldBind(&userRequest)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad_request",
			"message": err.Error(),
		})
		return
	}

	fmt.Println("User =>", &userRequest)
	successMessage, err2 := u.userService.Register(&userRequest)

	if err2 != nil {
		c.JSON(err2.Status(), gin.H{
			"error":   err2.Message(),
			"message": err2.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, successMessage)
}

// UpdateUser godoc
// @Schemes http
// @Tags Users
// @Summary Update user
// @Description Only can update username & email only
// @Param        RequestBody   body      dto.UpdateUserDataRequest  true  "User Request"
// @Accept json
// @Produce json
// @Success 200 {object} dto.UpdateUserDataResponse
// @Failure 400 {object} object
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFAYS5jb20iLCJleHAiOjE2NjYyODU5MzYsImlkIjoyfQ.ahlQ3Icb2h72Aam9jWgZFM-TultwjKzymM7DGKHVVKU)
// @Router /users [put]
func (u *userRestHandler) UpdateUserData(c *gin.Context) {
	var updateUserDataRequest dto.UpdateUserDataRequest
	var err error
	var userData entity.User

	contentType := helpers.GetContentType(c)
	if contentType == helpers.AppJSON {
		err = c.ShouldBindJSON(&updateUserDataRequest)
	} else {
		err = c.ShouldBind(&updateUserDataRequest)
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

	response, err2 := u.userService.UpdateUserData(userData.ID, &updateUserDataRequest)
	if err2 != nil {
		c.JSON(err2.Status(), gin.H{
			"error":   err2.Error(),
			"message": err2.Message(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteUser godoc
// @Schemes http
// @Tags Users
// @Summary Delete user
// @Description Delete User
// @Accept json
// @Produce json
// @Success 200 {object} dto.DeleteUserResponse
// @Failure 400 {object} object
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImJAYi5jb20iLCJleHAiOjE2NjYyODYxNTUsImlkIjo1fQ.VWqZ38H-gi87mwL3A75YloIBA27JqucEMZsR2TJtIBU)
// @Router /users [delete]
func (u *userRestHandler) DeleteUser(c *gin.Context) {
	var userData entity.User
	if value, ok := c.MustGet("userData").(entity.User); !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err_message": "unauthorized",
		})
		return
	} else {
		userData = value
	}

	response, err2 := u.userService.DeleteUser(userData.ID)
	if err2 != nil {
		c.JSON(err2.Status(), gin.H{
			"error":   err2.Error(),
			"message": err2.Message(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
