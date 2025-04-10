package main

func sortArrayByParity(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	maxNum := nums[0]
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] > maxNum {
			maxNum = nums[i]
		}
	}

	out := make([]int, len(nums))
	pos1 := -1
	pos2 := len(out)

	for i := range nums {
		if nums[i]%2 == 0 {
			pos1++
			out[pos1] = nums[i]
		} else {
			pos2--
			out[pos2] = nums[i]
		}
	}
	return out
}
