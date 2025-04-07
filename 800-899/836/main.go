package main

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	if rec1[2] <= rec2[0] || rec1[0] >= rec2[2] {
		return false
	}
	if rec1[3] <= rec2[1] || rec1[1] >= rec2[3] {
		return false
	}
	return true
}
