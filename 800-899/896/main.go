package main

import (
	"fmt"
)

func main() {
	fmt.Println(isMonotonic([]int{3, 3, 4}))
}

func isMonotonic(nums []int) bool {
	if len(nums) <= 2 {
		return true
	}
	typeOf := 0
	if nums[0] == nums[len(nums)-1] {
		typeOf = equality
	} else if nums[0] < nums[len(nums)-1] {
		typeOf = greater
	} else if nums[0] > nums[len(nums)-1] {
		typeOf = smaller
	}
	for i := 1; i < len(nums); i++ {
		if !check(nums[i-1], nums[i], typeOf) {
			return false
		}
	}
	return true
}

const (
	equality = iota
	greater
	smaller
)

func check(one, two, typeof int) bool {
	switch typeof {
	case equality:
		return one == two
	case greater:
		return two >= one
	case smaller:
		return two <= one
	}
	return false
}
