package main

func firstBadVersion(n int) int {
	minNum := 1
	maxNum := n

	for {
		if maxNum-minNum < 2 {
			if isBadVersion(minNum) {
				return minNum
			}
			return maxNum
		}
		center := minNum + (maxNum-minNum)/2
		if isBadVersion(center) {
			maxNum = center
		} else {
			minNum = center
		}
	}
}

func isBadVersion(version int) bool {
	if version > 10 {
		return true
	}
	return false
}
