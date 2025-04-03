package main

func pivotIndex(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var lSum, rSum int
	for i := 0; i < len(nums); i++ {
		rSum += nums[i]
	}

	for i := range nums {
		rSum -= nums[i]
		if rSum == lSum {
			return i
		}
		lSum += nums[i]
	}
	return -1
}
