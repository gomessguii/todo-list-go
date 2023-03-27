package main

import (
	"github.com/gin-gonic/gin"
	userCtrlPkg "github.com/gomessguii/todo-list-go/user/controller"
	userDmPkg "github.com/gomessguii/todo-list-go/user/domain"
	userRepoPkg "github.com/gomessguii/todo-list-go/user/repository"

	todoCtrlPkg "github.com/gomessguii/todo-list-go/todo/controller"
	todoDmPkg "github.com/gomessguii/todo-list-go/todo/domain"
	todoRepoPkg "github.com/gomessguii/todo-list-go/todo/repository"
)

func main() {
	userDm := userDmPkg.New(userRepoPkg.New())
	userCtrl := userCtrlPkg.New(userDm)

	todoDm := todoDmPkg.New(userDm, todoRepoPkg.New())
	todoCtrl := todoCtrlPkg.New(todoDm)

	r := gin.Default()
	r.POST("/login", userCtrl.Login)

	r.GET("/todo/:id", todoCtrl.GetTodo)
	r.POST("/todo/:id/complete", todoCtrl.CompleteTodo)
	r.POST("/todo/new", todoCtrl.CreateTodo)

	r.Run()
}
