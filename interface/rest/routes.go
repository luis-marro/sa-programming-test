package main

import (
	"github.com/gin-gonic/gin"
	"github.com/luis-marro/sa-programming-test/application"
	"github.com/luis-marro/sa-programming-test/domain/domain/entity"
	"net/http"
)

type intervalsRequest struct {
	Includes [][2]int `json:"includes"`
	Excludes [][2]int `json:"excludes"`
}

func runIntervals(c *gin.Context) {
	var requestBody intervalsRequest
	var includes []entity.Interval
	var excludes []entity.Interval

	intervalService := application.LocateIntervalsService()
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Build the intervals for the request
	for _, include := range requestBody.Includes {
		includes = append(includes, entity.Interval{
			Beginning: include[0],
			End:       include[1],
		})
	}
	for _, exclude := range requestBody.Excludes {
		excludes = append(excludes, entity.Interval{
			Beginning: exclude[0],
			End:       exclude[1],
		})
	}

	output := intervalService.ExcludeIntervals(includes, excludes)
	c.JSON(http.StatusOK, output)
}

func runFizzBuzz(c *gin.Context) {
	fizzbuzzService := application.LocateFizzBuzzService()

	c.JSON(http.StatusOK, fizzbuzzService.CallFizzBuzz())
}
