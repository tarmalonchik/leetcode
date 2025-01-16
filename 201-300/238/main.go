package main

func productExceptSelf(nums []int) []int {
	out := make([]int, len(nums))

	var prevVal = 1
	for i := range nums {
		out[i] = prevVal
		prevVal = nums[i] * prevVal
	}

	prevVal = 1
	for i := len(nums) - 1; i >= 0; i-- {
		out[i] = prevVal * out[i]
		prevVal = nums[i] * prevVal
	}
	return out
}
