package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"projectgo/model"
	"projectgo/service"
	"projectgo/utils"
)

type UserHandler struct {
	Service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (uh *UserHandler) Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{Status: http.StatusBadRequest, Message: "Invalid JSON Request"})
		return
	}

	if err := uh.Service.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, utils.Response{
		Status:  http.StatusCreated,
		Message: "User Created",
		Data:    map[string]interface{}{"user": user.UserID, "username": user.Username},
	})
}

func (uh *UserHandler) Login(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{
			Status: http.StatusBadRequest, Message: "Invalid JSON Request",
		})
		return
	}

	userData, err := uh.Service.GetUserByUsername(user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	accesstoken, err := utils.GenerateJWTAccessToken(userData.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed to generate access token",
		})
		return
	}

	refreshToken, err := utils.GenerateJWTRefreshToken(userData.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed to generate refresh token",
		})
		return
	}

	c.SetCookie("access_token", accesstoken, 5, "/", "", false, true)
	c.SetCookie("refresh_token", refreshToken, 3600, "/", "", false, true)

	c.JSON(http.StatusOK, utils.Response{
		Status:  http.StatusOK,
		Message: "User Logged In",
		Data: map[string]interface{}{
			"user_id": userData.UserID,
		},
	})
}

func (uh *UserHandler) Logout(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "", false, true)  // Mengatur cookie dengan durasi -1 untuk menghapusnya
	c.SetCookie("refresh_token", "", -1, "/", "", false, true) // Mengatur cookie dengan durasi -1 untuk menghapusnya

	c.JSON(http.StatusOK, utils.Response{
		Status:  http.StatusOK,
		Message: "User Logged Out",
	})
}

func (uh *UserHandler) GetAll(c *gin.Context) {
	users, err := uh.Service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed to get all users",
		})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    users,
	})
}
