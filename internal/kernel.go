package internal

import (
    "budget-tracker-api/internal/identity"
    "github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	identity.Init(r)
}