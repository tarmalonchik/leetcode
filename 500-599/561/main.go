package main

import (
	"sort"
)

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	sum := 0
	for i := 0; i < len(nums); i += 2 {
		sum += minNum(nums[i], nums[i+1])
	}
	return sum
}

func minNum(a, b int) int {
	if a < b {
		return a
	}
	return b
}
