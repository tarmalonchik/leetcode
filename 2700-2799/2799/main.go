package main

func countCompleteSubarrays(nums []int) int {
	out := 1
	mp := make(map[int]int)

	for i := range nums {
		mp[nums[i]]++
	}

	for {
		num := recursiveRight(nums, mp)
		out += num
		mp[nums[0]]--
		if mp[nums[0]] == 0 {
			break
		} else {
			out++
		}
		nums = nums[1:]
		if len(nums) == 0 {
			break
		}
	}
	return out
}

func recursiveRight(nums []int, mp map[int]int) int {
	out := 0
	if len(nums) == 0 {
		return 0
	}
	mp[nums[len(nums)-1]]--
	if mp[nums[len(nums)-1]] > 0 {
		out++
		out += recursiveRight(nums[:len(nums)-1], mp)
	}
	mp[nums[len(nums)-1]]++
	return out
}
