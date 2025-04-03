package main

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	pos1 := 0
	pos2 := len(nums) - 1

	for {
		if pos2-pos1 <= 1 {
			if nums[pos1] == target {
				return pos1
			}
			if nums[pos2] == target {
				return pos2
			}
			return -1
		}
		center := pos1 + (pos2-pos1)/2
		if nums[center] > target {
			pos2 = center
		} else {
			pos1 = center
		}
	}
}
