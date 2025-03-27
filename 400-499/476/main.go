package main

func findComplement(num int) int {
	n := 1
	for {
		if n > num {
			return n - 1 ^ num
		}
		n *= 2
	}
}
