package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gomessguii/todo-list-go/user/models"
)

type IUserController interface {
	Login(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	IsLoggedIn(ctx *gin.Context)
	Logout(ctx *gin.Context)
}

type userController struct {
	userDomain IUserDomain
}

func (u *userController) Login(ctx *gin.Context) {
	username := ctx.GetHeader("username")
	if username == "" {
		ctx.AbortWithError(400, fmt.Errorf("username is required"))
	}

	password := ctx.GetHeader("password")
	if password == "" {
		ctx.AbortWithError(400, fmt.Errorf("password is required"))
	}

	err := u.userDomain.Login(username, password)
	if err != nil {
		ctx.AbortWithError(403, err)
	}

	ctx.JSON(200, "logged in successfully")
}

func (ctrl *userController) CreateUser(ctx *gin.Context) {
	userRequest := &models.UserTransfer{}
	err := ctx.BindJSON(&userRequest)
	if err != nil {
		ctx.JSON(400, err)
	}
	err = ctrl.userDomain.CreateUser(userRequest.ToModel())
	if err != nil {
		ctx.JSON(400, err)
	}
	ctx.JSON(201, "user created")
}

func (ctrl *userController) IsLoggedIn(ctx *gin.Context) {
	ctx.JSON(200, ctrl.userDomain.IsLoggedIn())
}

func (ctrl *userController) Logout(ctx *gin.Context) {
	ctrl.userDomain.Logout()
	ctx.JSON(200, "logged out")
}

func New(userDomain IUserDomain) IUserController {
	return &userController{userDomain: userDomain}
}