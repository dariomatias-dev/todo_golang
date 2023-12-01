package routes

import (
	"github.com/gin-gonic/gin"

	db "github.com/darionatias-dev/todo_golang/api/db/sqlc"
)

func AppRoutes(router *gin.Engine, dbQueries *db.Queries) *gin.RouterGroup {
	v1 := router.Group("/v1")
	{
		todos := v1.Group("")
		{
			todos.POST("/todo")
			todos.GET("/todo/:id")
			todos.GET("/todos")
			todos.PATCH("/todo/:id")
			todos.PATCH("/todo-status/:id")
			todos.DELETE("/todo/:id")
		}
	}

	return v1
}
