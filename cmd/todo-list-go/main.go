package main

import (
	"log"

	"github.com/gin-gonic/gin"
	userCtrlPkg "github.com/gomessguii/todo-list-go/user/controller"
	userDmPkg "github.com/gomessguii/todo-list-go/user/domain"
	userModels "github.com/gomessguii/todo-list-go/user/models"
	userRepoPkg "github.com/gomessguii/todo-list-go/user/repository"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	todoCtrlPkg "github.com/gomessguii/todo-list-go/todo/controller"
	todoDmPkg "github.com/gomessguii/todo-list-go/todo/domain"
	todoModels "github.com/gomessguii/todo-list-go/todo/models"
	todoRepoPkg "github.com/gomessguii/todo-list-go/todo/repository"
)

func main() {
	db, err := gorm.Open(sqlite.Open("databasse.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&todoModels.TodoTransfer{}, &userModels.UserTransfer{})

	userDm := userDmPkg.New(userRepoPkg.New(db))
	userCtrl := userCtrlPkg.New(userDm)

	todoDm := todoDmPkg.New(userDm, todoRepoPkg.New(db))
	todoCtrl := todoCtrlPkg.New(todoDm)

	r := gin.Default()
	r.POST("/login", userCtrl.Login)
	r.POST("/logout", userCtrl.Logout)

	r.POST("/user/new", userCtrl.CreateUser)
	r.POST("/user/loggedin", userCtrl.IsLoggedIn)

	r.GET("/todo/:id", todoCtrl.GetTodo)
	r.POST("/todo/:id/complete", todoCtrl.CompleteTodo)
	r.POST("/todo/new", todoCtrl.CreateTodo)

	log.Println("server listening on :8080")
	r.Run()
}
