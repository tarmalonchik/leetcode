package main

func moveZeroes(nums []int) {
	zeroPos := -1
	numPos := -1

	numPos = find(nums, 0, false)
	if numPos == -1 {
		return
	}

	zeroPos = find(nums, 0, true)
	if zeroPos == -1 {
		return
	}

	for {
		if zeroPos < numPos {
			nums[zeroPos] = nums[numPos]
			nums[numPos] = 0
		} else {
			numPos = find(nums, numPos+1, false)
			if numPos == -1 {
				break
			}
			continue
		}
		numPos = find(nums, numPos+1, false)
		if numPos == -1 {
			break
		}
		zeroPos = find(nums, zeroPos+1, true)
		if zeroPos == -1 {
			break
		}
	}
}

func find(nums []int, startIdx int, isZero bool) int {
	out := -1
	for i := startIdx; i < len(nums); i++ {
		if isZero && nums[i] == 0 {
			return i
		}
		if !isZero && nums[i] != 0 {
			return i
		}
	}
	return out
}
