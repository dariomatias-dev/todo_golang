package middleware

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func HandleError() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if pgErr, ok := err.(*pq.Error); ok {
					handleDatabaseError(ctx, pgErr)
					return
				} else if err == sql.ErrNoRows {
					ctx.AbortWithStatusJSON(http.StatusNotFound, nil)
					return
				}

				ctx.AbortWithStatusJSON(
					http.StatusInternalServerError,
					gin.H{
						"error":   "Internal server error",
						"message": err,
					},
				)
				return
			}
		}()

		ctx.Next()
	}
}

func handleDatabaseError(ctx *gin.Context, pgErr *pq.Error) {
	switch pgErr.Code {
	case "23505":
		fieldName := extractFieldName(pgErr.Message)
		errorMessage := fmt.Sprintf("Value of the %s field already exists", *fieldName)

		ctx.AbortWithStatusJSON(
			http.StatusConflict,
			gin.H{
				"message":   errorMessage,
				"fieldName": fieldName,
			},
		)
		return
	}
}

func extractFieldName(message string) *string {
	char := "_"

	startIndex := strings.Index(message, char)
	endIndex := strings.LastIndex(message, char)

	fieldName := ""

	if startIndex == -1 || endIndex == -1 {
		return &fieldName
	}

	fieldName = message[startIndex+1 : endIndex]

	return &fieldName
}
