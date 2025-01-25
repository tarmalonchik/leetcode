package main

func guessNumber(n int) int {
	minNum := 0
	maxNum := n
	for {
		if maxNum-minNum <= 1 {
			if guess(minNum) == 0 {
				return minNum
			}
			if guess(maxNum) == 0 {
				return maxNum
			}
			return 0
		}
		center := minNum + (maxNum-minNum)/2
		if guess(center) < 0 {
			maxNum = center
			continue
		}
		if guess(center) > 0 {
			minNum = center
			continue
		}
		return center
	}
}

func guess(num int) int {
	pick := 6
	if num > pick {
		return -1
	}
	if num == pick {
		return 0
	}
	return 1
}
