package main

import (
    "budget-tracker-api/internal"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	internal.Init(r)

	err := r.Run(":8000")
	if err != nil {
		return
	}
}