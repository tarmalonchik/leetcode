package main

func findLengthOfLCIS(nums []int) int {
	maxN := 0
	n := 0
	prevNum := 0
	for i := range nums {
		if n == 0 {
			n++
			prevNum = nums[i]
			continue
		} else if nums[i] > prevNum {
			prevNum = nums[i]
			n++
			continue
		} else {
			if n > maxN {
				maxN = n
			}
			n = 1
			prevNum = nums[i]
		}
	}
	if n > maxN {
		maxN = n
	}
	return maxN
}
