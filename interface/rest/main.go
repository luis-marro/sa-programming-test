package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.POST("/algorithms/intervals", runIntervals)

	_ = r.Run()
}
