package main

func isPerfectSquare(num int) bool {
	if num <= 0 {
		return false
	}

	minNum := 0
	maxNum := num

	for {
		if maxNum-minNum <= 1 {
			if minNum*minNum == num {
				return true
			}
			if maxNum*maxNum == num {
				return true
			}
			return false
		}
		center := minNum + (maxNum-minNum)/2
		if center*center > num {
			maxNum = center
		} else {
			minNum = center
		}
	}
}
