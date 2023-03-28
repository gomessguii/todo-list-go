package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gomessguii/todo-list-go/todo/models"
)

type ITodoController interface {
	CreateTodo(ctx *gin.Context)
	CompleteTodo(ctx *gin.Context)
	GetTodo(ctx *gin.Context)
}

type todoController struct {
	todoDomain ITodoDomain
}

func New(todoDomain ITodoDomain) ITodoController {
	return &todoController{todoDomain}
}

func (ctrl *todoController) CompleteTodo(ctx *gin.Context) {
	idStr, present := ctx.Params.Get("id")
	if !present {
		ctx.JSON(400, "id not informed")
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(400, "invalid id")
	}
	err = ctrl.todoDomain.CompleteTodo(id)
	if err != nil {
		ctx.JSON(400, err)
	}
	ctx.JSON(200, "Todo completed successfully")
}

func (ctrl *todoController) CreateTodo(ctx *gin.Context) {
	todoRequest := &models.TodoTransfer{}
	err := ctx.BindJSON(&todoRequest)
	if err != nil {
		ctx.JSON(400, err)
	}
	err = ctrl.todoDomain.CreateTodo(todoRequest.ToModel())
	if err != nil {
		ctx.JSON(400, err)
	}
	ctx.JSON(201, "todo created")
}

func (ctrl *todoController) GetTodo(ctx *gin.Context) {
	idStr, present := ctx.Params.Get("id")
	if !present {
		ctx.JSON(400, "id not informed")
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(400, "invalid id")
	}
	todo, err := ctrl.todoDomain.GetTodo(id)
	if err != nil {
		ctx.JSON(400, err)
	}
	ctx.JSON(200, todo.ToTransfer())
}
