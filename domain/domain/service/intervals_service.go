package service

import "sa-programming-test/domain/domain/entity"

type CalculateIntervalExclude interface {
	ExcludeIntervals(includes []entity.Interval, excludes []entity.Interval) []entity.Interval
}
