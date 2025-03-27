package main

func hammingDistance(x int, y int) int {
	out := 0
	xor := x ^ y

	for {
		if xor%2 == 1 {
			out++
		}
		xor /= 2
		if xor == 0 {
			return out
		}
	}
}
