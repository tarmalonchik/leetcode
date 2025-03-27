package main

func bitwiseComplement(n int) int {
	if n == 0 {
		return 1
	}
	num := 1
	for {
		if num > n {
			return num - 1 ^ n
		}
		num *= 2
	}
}
