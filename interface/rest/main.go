package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/luis-marro/sa-programming-test/interface/rest/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title			Solutions Architect Programming Challenge
// @description	API containing the two challenges pertaining to the SA programming challenge
// @version		1.0
func main() {
	r := gin.Default()

	algorithms := r.Group("/algorithm")
	algorithms.Use(CORSMiddleware())
	algorithms.POST("/intervals", runIntervals)
	algorithms.GET("/fizzbuzz", runFizzBuzz)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	_ = r.Run()
}
