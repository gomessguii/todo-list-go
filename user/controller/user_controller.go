package controller

import "github.com/gin-gonic/gin"

type IUserController interface {
	Login(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	IsLoggedIn(ctx *gin.Context)
	Logout(ctx *gin.Context)
}

func New(IUserDomain) IUserController {
	panic("implement me")
}