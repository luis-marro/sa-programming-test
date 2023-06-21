package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.POST("/algorithms/intervals", runIntervals)
	r.GET("/algorithms/fizzbuzz", runFizzBuzz)

	_ = r.Run()
}
