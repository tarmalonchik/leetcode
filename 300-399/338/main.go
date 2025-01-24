package main

func countBits(n int) []int {
	out := make([]int, 0, n+1)
	offset := 1
	numPow := 2
	for i := 0; i <= n; i++ {
		if i >= numPow {
			numPow *= 2
			offset++
		}
		out = append(out, getBitsNum(i, offset))
	}
	return out
}

func getBitsNum(n int, maxBits int) int {
	out := 0
	for i := 0; i < maxBits; i++ {
		if getBit(uint32(n), uint8(i)) {
			out += 1
		}
	}
	return out
}

func getBit(num uint32, pos uint8) bool {
	if num&(1<<pos) > 0 {
		return true
	}
	return false
}
