package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/darionatias-dev/todo_golang/api/controllers"
	db "github.com/darionatias-dev/todo_golang/api/db/sqlc"
)

func AppRoutes(router *gin.Engine, dbQueries *db.Queries) *gin.RouterGroup {
	todoController := controllers.NewTodoController(dbQueries)

	v1 := router.Group("/v1")
	{
		todos := v1.Group("")
		{
			todos.POST("/todo", todoController.Create)
			todos.GET("/todo/:id", todoController.GetOne)
			todos.GET("/todos", todoController.GetAll)
			todos.PATCH("/todo/:id", todoController.Update)
			todos.PATCH("/todo-status/:id", todoController.UpdateStatus)
			todos.DELETE("/todo/:id", todoController.Delete)
		}
	}

	return v1
}
