package main

import (
	"strconv"
)

func fizzBuzz(n int) []string {
	out := make([]string, n)
	three := 0
	five := 0
	for i := 1; i <= n; i++ {
		three++
		five++
		if three == 3 && five == 5 {
			out[i-1] = "FizzBuzz"
			three = 0
			five = 0
		} else if three == 3 {
			out[i-1] = "Fizz"
			three = 0
		} else if five == 5 {
			out[i-1] = "Buzz"
			five = 0
		} else {
			out[i-1] = strconv.Itoa(i)
		}
	}
	return out
}
