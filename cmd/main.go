package main

import (
    "budget-tracker-api/internal"
    "budget-tracker-api/internal/config"
    "budget-tracker-api/internal/database"
    "github.com/gin-gonic/gin"
    "log"
)

func main() {
	var err error
	gin.SetMode(gin.DebugMode)

    err = config.LoadConfig()
    if err != nil {
		log.Fatal(err)
        return
    }
	database.Connect()

	router := gin.Default()
	internal.Init(router)

	err = router.Run(":8000")
	if err != nil {
		return
	}
}