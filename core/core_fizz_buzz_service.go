package core

import (
	"sa-programming-test/domain/domain/service"
)

type CoreFizzBuzzService struct {
}

func (c CoreFizzBuzzService) CallFizzBuzz() []interface{} {
	output := make([]interface{}, service.FizzBuzzUpperLimit-service.FizzBuzzLowerLimit+1)

	for i := 1; i <= len(output); i++ {
		if i%15 == 0 {
			output[i-1] = service.FizzBuzzString
		} else if i%5 == 0 {
			output[i-1] = service.BuzzString
		} else if i%3 == 0 {
			output[i-1] = service.FizzString
		} else {
			output[i-1] = i
		}
	}

	return output
}
