package main

func intersect(nums1 []int, nums2 []int) []int {
	var (
		mp  = make(map[int]int)
		out = make([]int, 0)
	)

	for i := range nums1 {
		mp[nums1[i]]++
	}

	for i := range nums2 {
		val, ok := mp[nums2[i]]
		if ok && val > 0 {
			mp[nums2[i]]--
			out = append(out, nums2[i])
		}
	}
	return out
}
