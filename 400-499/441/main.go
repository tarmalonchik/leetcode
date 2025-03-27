package main

func arrangeCoins(n int) int {
	var out = 0
	var counter = 1
	for {
		n -= counter
		if n < 0 {
			break
		}
		out++
		counter++
	}
	return out
}
