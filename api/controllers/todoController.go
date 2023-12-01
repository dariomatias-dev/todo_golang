package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	db "github.com/dariomatias-dev/todo_golang/api/db/sqlc"
	"github.com/dariomatias-dev/todo_golang/api/models"
	"github.com/dariomatias-dev/todo_golang/api/utils"
)

type todoController struct {
	DbQueries *db.Queries
}

func NewTodoController(dbQueries *db.Queries) *todoController {
	return &todoController{
		DbQueries: dbQueries,
	}
}

func (tc *todoController) Create(ctx *gin.Context) {
	data := models.CreateTodo{}

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ID := uuid.NewString()
	arg := db.CreateTodoParams{
		ID:          ID,
		Title:       data.Title,
		Description: data.Description,
		Status:      false,
	}

	todo, err := tc.DbQueries.CreateTodo(ctx, arg)

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusCreated, todo)
}

func (tc *todoController) FindOne(ctx *gin.Context) {
	ID := ctx.Param("id")

	todo, err := tc.DbQueries.GetTodo(ctx, ID)

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, todo)
}

func (tc *todoController) FindAll(ctx *gin.Context) {
	todos, err := tc.DbQueries.GetTodos(ctx)

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, todos)
}

func (tc *todoController) UpdateStatus(ctx *gin.Context) {
	ID := ctx.Param("id")

	data := models.UpdateTodoStatus{}

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	arg := db.UpdateTodoStatusParams{
		ID:     ID,
		Status: *data.Status,
	}

	todo, err := tc.DbQueries.UpdateTodoStatus(ctx, arg)

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, todo)
}

func (tc *todoController) Update(ctx *gin.Context) {
	ID := ctx.Param("id")

	data := models.UpdateTodo{}

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	getValue := utils.GetValue{}

	arg := db.UpdateTodoParams{
		ID:          ID,
		Title:       *getValue.NullString(data.Title),
		Description: *getValue.NullString(data.Description),
		Status:      *getValue.NullBool(nil),
	}

	todo, err := tc.DbQueries.UpdateTodo(ctx, arg)

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, todo)
}

func (tc *todoController) Delete(ctx *gin.Context) {
	ID := ctx.Param("id")

	todo, err := tc.DbQueries.DeleteTodo(ctx, ID)

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, todo)
}
