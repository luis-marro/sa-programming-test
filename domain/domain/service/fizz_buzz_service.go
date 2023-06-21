package service

type FizzBuzzService interface {
	CallFizzBuzz() []interface{}
}

const FizzBuzzLowerLimit = 1
const FizzBuzzUpperLimit = 100

const FizzString = "Fizz"
const BuzzString = "Buzz"
const FizzBuzzString = "FizzBuzz"
