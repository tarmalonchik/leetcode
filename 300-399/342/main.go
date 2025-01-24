package main

func isPowerOfFour(n int) bool {
	if n == 1 {
		return true
	}
	if n < 4 {
		return false
	}

	num := 4
	for {
		if n == num {
			return true
		}
		if num > n {
			return false
		}
		num *= 4
	}
}
