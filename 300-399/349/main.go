package main

func intersection(nums1 []int, nums2 []int) []int {
	var lw []int
	var bg []int
	if len(nums1) <= len(nums2) {
		lw = nums1
		bg = nums2
	} else {
		lw = nums2
		bg = nums1
	}

	mp := make(map[int]interface{})

	for i := range lw {
		mp[lw[i]] = nil
	}

	out := make([]int, 0)

	for i := range bg {
		if _, ok := mp[bg[i]]; ok {
			delete(mp, bg[i])
			out = append(out, bg[i])
		}
	}
	return out
}
