package main

func findMin(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	if nums[0] < nums[len(nums)-1] {
		return nums[0]
	}

	minPos := 0
	maxPos := len(nums) - 1

	for {
		if maxPos-minPos <= 1 {
			return minNum(nums[maxPos], nums[minPos])
		}
		center := minPos + ((maxPos - minPos) / 2)
		if nums[center] < nums[minPos] {
			maxPos = center
		} else {
			minPos = center
		}
	}
}

func minNum(a, b int) int {
	if a < b {
		return a
	}
	return b
}
