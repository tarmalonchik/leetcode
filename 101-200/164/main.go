package main

import (
	"math"
)

// Radix sort
// https://ru.wikipedia.org/wiki/Поразрядная_сортировка

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	sort(nums)
	maxGap := 0

	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] > maxGap {
			maxGap = nums[i] - nums[i-1]
		}
	}
	return maxGap
}

func sort(nums []int) {
	var idx = 0
	for {
		sortedNums := make([][]int, 10)
		needBreak := true
		for i := range nums {
			sortedNumsIdx, valid := getNumInSpecificIdx(nums[i], idx)
			if valid {
				needBreak = false
			}
			sortedNums[sortedNumsIdx] = append(sortedNums[sortedNumsIdx], nums[i])
		}
		if needBreak {
			break
		}

		position := 0
		for i := range sortedNums {
			for j := range sortedNums[i] {
				nums[position] = sortedNums[i][j]
				position++
			}
		}
		idx++
	}
}

func getNumInSpecificIdx(num, idx int) (int, bool) {
	divide := int(math.Pow10(idx))
	num = num / divide
	if num > 0 {
		return num % 10, true
	}
	return 0, false
}
