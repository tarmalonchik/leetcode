package main

func containsDuplicate(nums []int) bool {
	mp := make(map[int]interface{}, len(nums))
	for i := range nums {
		_, ok := mp[nums[i]]
		if ok {
			return true
		}
		mp[nums[i]] = nil
	}
	return false
}
