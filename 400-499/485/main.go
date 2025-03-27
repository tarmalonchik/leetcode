package main

func findMaxConsecutiveOnes(nums []int) int {
	maxOut := 0
	out := 0
	for i := range nums {
		if nums[i] == 1 {
			out++
			continue
		}
		if out > maxOut {
			maxOut = out
		}
		out = 0
	}
	if out > maxOut {
		return out
	}
	return maxOut
}
