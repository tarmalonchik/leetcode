package main

func findShortestSubArray(nums []int) int {
	mp := make(map[int]int)
	degree := 0
	for i := range nums {
		mp[nums[i]]++
		val := mp[nums[i]]
		if val > degree {
			degree = val
		}
	}
	if degree == 1 {
		return 1
	}

	for key, val := range mp {
		if val != degree {
			delete(mp, key)
		}
	}
	return findRecursion(nums, mp)
}

func findRecursion(nums []int, mp map[int]int) int {
	pos1 := 0
	pos2 := len(nums) - 1

	for {
		if pos1 >= pos2 {
			return 1
		}

		if _, ok := mp[nums[pos1]]; !ok {
			pos1++
			continue
		}
		if _, ok := mp[nums[pos2]]; !ok {
			pos2--
			continue
		}

		delete(mp, nums[pos1])
		if len(mp) == 0 {
			mp[nums[pos1]] = 1
			return pos2 - pos1 + 1
		}
		a := findRecursion(nums[pos1+1:pos2+1], mp)
		mp[nums[pos1]] = 1

		delete(mp, nums[pos2])
		if len(mp) == 0 {
			mp[nums[pos2]] = 1
			return pos2 - pos1 + 1
		}
		b := findRecursion(nums[pos1:pos2], mp)
		mp[nums[pos2]] = 1

		if a < b {
			return a
		}
		return b
	}
}
