package controller

import "github.com/gin-gonic/gin"

type ITodoController interface {
	CreateTodo(ctx *gin.Context)
	CompleteTodo(ctx *gin.Context)
	GetTodo(ctx *gin.Context)
}

func New(ITodoDomain) ITodoController {
	panic("implement me")
}