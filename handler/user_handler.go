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
		c.JSON(http.StatusBadRequest, utils.Response{Status: http.StatusBadRequest, Message: "Invalid JSON Request"})
		return
	}

	userData, err := uh.Service.GetUserByUsername(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Status:  http.StatusOK,
		Message: "User Logged In",
		Data:    map[string]interface{}{"user_id": userData.UserID, "username": userData.Username},
	})
}
