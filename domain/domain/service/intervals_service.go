package service

import "github.com/luis-marro/sa-programming-test/domain/domain/entity"

type DomainIntervalsService interface {
	ExcludeIntervals(includes []entity.Interval, excludes []entity.Interval) []entity.Interval
}
