package core

import (
	"reflect"
	"sa-programming-test/domain/domain/entity"
	"testing"
)

type excludeIntervalTestTable struct {
	includes []entity.Interval
	excludes []entity.Interval
	expects  []entity.Interval
}

var testExcludeIntervalTable = []excludeIntervalTestTable{
	{
		includes: []entity.Interval{
			{
				Beginning: 10,
				End:       100,
			},
			{
				Beginning: 200,
				End:       300,
			},
			{
				Beginning: 400,
				End:       500,
			},
		},
		excludes: []entity.Interval{
			{
				Beginning: 95,
				End:       205,
			},
			{
				Beginning: 410,
				End:       420,
			},
		},
		expects: []entity.Interval{
			{
				Beginning: 10,
				End:       94,
			},
			{
				Beginning: 206,
				End:       300,
			},
			{
				Beginning: 400,
				End:       409,
			},
			{
				Beginning: 421,
				End:       500,
			},
		},
	},
	{
		includes: []entity.Interval{
			{
				Beginning: 10,
				End:       100,
			},
		},
		excludes: []entity.Interval{
			{
				Beginning: 20,
				End:       30,
			},
		},
		expects: []entity.Interval{
			{
				Beginning: 10,
				End:       19,
			},
			{
				Beginning: 31,
				End:       100,
			},
		},
	},
	{
		includes: []entity.Interval{
			{
				Beginning: 50,
				End:       5000,
			},
			{
				Beginning: 10,
				End:       100,
			},
		},
		excludes: []entity.Interval{},
		expects: []entity.Interval{
			{
				Beginning: 10,
				End:       5000,
			},
		},
	},
	{
		includes: []entity.Interval{
			{
				Beginning: 10,
				End:       100,
			},
			{
				Beginning: 200,
				End:       300,
			},
		},
		excludes: []entity.Interval{
			{
				Beginning: 95,
				End:       205,
			},
		},
		expects: []entity.Interval{
			{
				Beginning: 10,
				End:       94,
			},
			{
				Beginning: 206,
				End:       300,
			},
		},
	},
	{
		includes: []entity.Interval{},
		excludes: []entity.Interval{
			{
				Beginning: 95,
				End:       205,
			},
		},
		expects: []entity.Interval{},
	},
}

func TestCoreIntervalsService_ExcludeIntervals(t *testing.T) {
	c := CoreIntervalsService{}

	for _, test := range testExcludeIntervalTable {
		got := c.ExcludeIntervals(test.includes, test.excludes)
		if !reflect.DeepEqual(got, test.expects) {
			t.Errorf("Output %+v does not match expected value %+v", got, test.expects)
		}
	}
}
