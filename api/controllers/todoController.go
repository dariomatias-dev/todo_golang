package controllers

import (
	"github.com/gin-gonic/gin"

	db "github.com/dariomatias-dev/todo_golang/api/db/sqlc"
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

}
func (tc *todoController) GetOne(ctx *gin.Context) {

}
func (tc *todoController) GetAll(ctx *gin.Context) {

}
func (tc *todoController) Update(ctx *gin.Context) {

}
func (tc *todoController) UpdateStatus(ctx *gin.Context) {

}
func (tc *todoController) Delete(ctx *gin.Context) {

}
