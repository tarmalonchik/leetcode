package main

import (
	"sort"
)

func distributeCandies(candyType []int) int {
	if len(candyType) == 0 {
		return 0
	}

	sort.Ints(candyType)
	maxCandies := len(candyType) / 2
	typeCounter := 1
	prevType := candyType[0]
	for i := 1; i < len(candyType); i++ {
		if prevType != candyType[i] {
			typeCounter++
			if typeCounter >= maxCandies {
				return maxCandies
			}
			prevType = candyType[i]
			continue
		}
	}
	return typeCounter
}
