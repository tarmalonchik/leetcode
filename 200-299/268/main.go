package main

func missingNumber(nums []int) int {
	checkerArr := make([]bool, len(nums)+1)
	for i := range nums {
		checkerArr[nums[i]] = true
	}
	for i := range checkerArr {
		if !checkerArr[i] {
			return i
		}
	}
	return 0
}
