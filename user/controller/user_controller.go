package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
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

func (u *userController) CreateUser(ctx *gin.Context) {
	panic("not implemented")
}

func (u *userController) IsLoggedIn(ctx *gin.Context) {
	panic("not implemented")
}

func (u *userController) Logout(ctx *gin.Context) {
	panic("not implemented")
}

func New(userDomain IUserDomain) IUserController {
	return &userController{userDomain: userDomain}
}