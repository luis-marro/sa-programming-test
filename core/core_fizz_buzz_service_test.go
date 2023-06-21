package core

import "testing"

type fizzBuzzTest struct {
	index int
	value interface{}
}

var fizzBuzzTestCases = []fizzBuzzTest{
	{
		index: 2,
		value: "Fizz",
	},
	{
		index: 9,
		value: "Buzz",
	},
	{
		index: 14,
		value: "FizzBuzz",
	},
	{
		index: 0,
		value: 1,
	},
	{
		index: 98,
		value: "Fizz",
	},
	{
		index: 99,
		value: "Buzz",
	},
}

func TestCallFizzBuzz(t *testing.T) {
	got := CoreFizzBuzzService{}.CallFizzBuzz()
	arrayLength := 100

	if len(got) != arrayLength { // Hardcoded value for testing purposes, should refactor test in constants are updated frequently/become dynamic.
		t.Errorf("Incorrect Array lenght. expected %v, got %v", arrayLength, len(got))
	}

	for _, value := range fizzBuzzTestCases {
		if got[value.index] != value.value {
			t.Errorf("Expected value at position %v to be %v, but got %v", value.index+1, got[value.index], value.value)
		}
	}
}
