package main

func findLHS(nums []int) int {
	maxNum := 0
	mp := make(map[int]int, len(nums))
	for i := range nums {
		mp[nums[i]]++
		val := mp[nums[i]]
		rVal := mp[nums[i]+1]
		lVal := mp[nums[i]-1]
		if rVal > 0 {
			maxNum = setMax(maxNum, val+rVal)
		}
		if lVal > 0 {
			maxNum = setMax(maxNum, val+lVal)
		}
	}
	return maxNum
}

func setMax(prevMax, newVal int) int {
	if newVal > prevMax {
		return newVal
	}
	return prevMax
}
