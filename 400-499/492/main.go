package main

import (
	"math"
)

func constructRectangle(area int) []int {
	start := int(math.Sqrt(float64(area)))
	for {
		if area%start == 0 {
			return maxFirst(start, area/start)
		}
		start++
	}
}

func maxFirst(a, b int) []int {
	if a > b {
		return []int{a, b}
	}
	return []int{b, a}
}
