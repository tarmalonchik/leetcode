package main

func findDisappearedNumbers(nums []int) []int {
	nums2 := make([]int, len(nums))
	out := make([]int, 0)

	for i := range nums {
		nums2[nums[i]-1] = nums[i]
	}

	for i := range nums {
		if nums2[i]-1 != i {
			out = append(out, i+1)
		}
	}
	return out
}
