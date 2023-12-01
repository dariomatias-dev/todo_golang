package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ValidUUIDMiddlware(ctx *gin.Context) {
	value := ctx.Param("id")

	ID, err := uuid.Parse(value)
	if err != nil || ID.Version() != 4 {
		errorMessage := fmt.Sprintf("Invalid ID %s", value)
		ctx.AbortWithStatusJSON(
			http.StatusBadGateway,
			gin.H{
				"message": errorMessage,
			},
		)
		return
	}
}
