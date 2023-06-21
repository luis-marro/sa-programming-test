package application

import (
	"github.com/luis-marro/sa-programming-test/core"
	"github.com/luis-marro/sa-programming-test/domain/domain/service"
)

func LocateIntervalsService() service.DomainIntervalsService {
	return core.CoreIntervalsService{}
}

func LocateFizzBuzzService() service.FizzBuzzService {
	return core.CoreFizzBuzzService{}
}
