package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/dariomatias-dev/todo_golang/api/controllers"
	db "github.com/dariomatias-dev/todo_golang/api/db/sqlc"
	"github.com/dariomatias-dev/todo_golang/api/middlewares"
)

func AppRoutes(router *gin.Engine, dbQueries *db.Queries) *gin.RouterGroup {
	todoController := controllers.NewTodoController(dbQueries)

	router.Use(middleware.HandleError())

	v1 := router.Group("/v1")
	{
		todos := v1.Group("")
		{
			todos.POST(
				"/todo",
				todoController.Create,
			)
			todos.GET(
				"/todo/:id",
				middleware.ValidUUIDMiddlware,
				todoController.FindOne,
			)
			todos.GET(
				"/todos",
				todoController.FindAll,
			)
			todos.PATCH(
				"/todo-status/:id",
				middleware.ValidUUIDMiddlware,
				todoController.UpdateStatus,
			)
			todos.PATCH(
				"/todo/:id",
				middleware.ValidUUIDMiddlware,
				todoController.Update,
			)
			todos.DELETE(
				"/todo/:id",
				middleware.ValidUUIDMiddlware,
				todoController.Delete,
			)
		}
	}

	return v1
}
