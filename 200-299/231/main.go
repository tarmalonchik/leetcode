package main

func isPowerOfTwo(n int) bool {
	if n == 1 || n == 2 {
		return true
	}
	if n <= 2 {
		return false
	}
	item := 2
	for {
		item = item * 2
		if item == n {
			return true
		}
		if item > n {
			return false
		}
	}
}
