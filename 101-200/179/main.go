package main

import (
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	out := ""
	sort.Slice(nums, func(a int, b int) bool {
		return compare(nums[b], nums[a])
	})
	for i := range nums {
		if (len(out) == 0 || out[len(out)-1] == '0') && nums[i] == 0 && i != len(nums)-1 {
			continue
		}
		out += strconv.Itoa(nums[i])
	}
	return out
}

func compare(a, b int) bool {
	aar := generateArr(a)
	bar := generateArr(b)

	aLen := len(aar)
	bLen := len(bar)

	isA := false
	if len(aar) < len(bar) {
		isA = true
		diff := len(bar) - len(aar)
		prefix := make([]int, 0)
		for i := len(bar) - diff; i < len(bar); i++ {
			prefix = append(prefix, bar[i])
		}
		aar = append(prefix, aar...)
	} else if len(bar) < len(aar) {
		diff := len(aar) - len(bar)
		prefix := make([]int, 0)
		for i := len(aar) - diff; i < len(aar); i++ {
			prefix = append(prefix, aar[i])
		}
		bar = append(prefix, bar...)
	}

	for i := len(aar) - 1; i >= 0; i-- {
		if aar[i] < bar[i] {
			return true
		} else if bar[i] < aar[i] {
			return false
		}
	}

	if isA {
		aar = aar[len(aar)-aLen:]
	} else {
		bar = bar[len(bar)-bLen:]
	}

	minNum := findMin(len(aar), len(bar)) - 1
	for i := minNum; i >= 0; i-- {
		if aar[i] < bar[i] {
			return false
		} else if aar[i] > bar[i] {
			return true
		}
	}
	return false
}

func findMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func generateArr(num int) []int {
	out := make([]int, 0)
	for {
		app := num % 10
		out = append(out, app)
		num /= 10
		if num == 0 {
			break
		}
	}
	return out
}
