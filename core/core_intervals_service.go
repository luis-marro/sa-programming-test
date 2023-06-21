package core

import (
	"sa-programming-test/domain/domain/entity"
	"sort"
)

// CoreIntervalsService creating this type of struct allows to pass other arguments to be injected in the object.
// this would before example a repository implementation of an interface, so a formal implementation will be injected at runtime
type CoreIntervalsService struct {
}

// ExcludeIntervals Algorithmic BigO Notation: O(n log n)
func (c CoreIntervalsService) ExcludeIntervals(includes []entity.Interval, excludes []entity.Interval) []entity.Interval {
	valuesMap := make(map[int]bool)
	var numericLine []int

	// Iterate over the intervals to feed the map and slice
	for _, value := range includes {
		valuesMap[value.Beginning] = true
		valuesMap[value.End] = true
		numericLine = append(numericLine, value.Beginning, value.End)
	}
	for _, value := range excludes {
		valuesMap[value.Beginning] = false
		valuesMap[value.End] = false
		numericLine = append(numericLine, value.Beginning, value.End)
	}

	// Sort the numeric line
	sort.Ints(numericLine)

	if len(includes) == 0 {
		return []entity.Interval{}
	} else if len(excludes) == 0 {
		return []entity.Interval{
			{
				Beginning: numericLine[0],
				End:       numericLine[len(numericLine)-1],
			},
		}
	}

	var output []entity.Interval
	var pivotInterval = entity.Interval{
		Beginning: numericLine[0],
		End:       numericLine[0],
	}

	for i := 1; i < len(numericLine); i++ {
		if !valuesMap[numericLine[i]] {
			pivotInterval.End = numericLine[i] - 1
			output = append(output, pivotInterval)
			if i >= len(numericLine) {
				break
			}
			i = i + getNextIntervalEnd(numericLine[i+1:], valuesMap)
			pivotInterval = entity.Interval{
				Beginning: numericLine[i] + 1,
				End:       numericLine[i] + 1,
			}
		} else {
			pivotInterval.End = numericLine[i]
			output = append(output, pivotInterval)
			if i+1 >= len(numericLine) {
				break
			}
			pivotInterval = entity.Interval{
				Beginning: numericLine[i+1],
				End:       numericLine[i+1],
			}
			i++ // In order to adjust the above value
		}
	}

	return output
}

// Helper function that returns the index of the number were the interval will stop
func getNextIntervalEnd(numericLine []int, valuesMap map[int]bool) int {
	for i, v := range numericLine {
		if !valuesMap[v] {
			return i + 1
		}
	}
	return len(numericLine) - 1
}
