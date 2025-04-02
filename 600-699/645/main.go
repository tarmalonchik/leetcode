package main

func findErrorNums(nums []int) []int {
	nums2 := make([]int, len(nums))
	out := make([]int, 0)

	for i := range nums {
		if nums2[nums[i]-1] != 0 {
			out = append(out, nums[i])
		}
		nums2[nums[i]-1] = nums[i]
	}

	for i := range nums2 {
		if nums2[i] == 0 {
			out = append(out, i+1)
		}
	}
	return out
}
