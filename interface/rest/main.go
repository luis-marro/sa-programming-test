package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	r := gin.Default()

	algorithms := r.Group("/algorithms")
	algorithms.Use(apiKeyAuthMiddleware())
	algorithms.POST("/intervals", runIntervals)
	algorithms.GET("/fizzbuzz", runFizzBuzz)

	_ = r.Run()
}
