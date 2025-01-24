package main

func addDigits(num int) int {
	if num < 10 {
		return num
	}
	for {
		newNum := 0
		for {
			if num < 10 {
				newNum += num
				break
			}
			newNum += num % 10
			num = num / 10
		}
		num = newNum
		if num < 10 {
			return num
		}
	}
}
