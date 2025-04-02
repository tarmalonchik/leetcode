package main

func hasAlternatingBits(n int) bool {
	maxPos := 1
	num := 1
	for {
		num *= 2
		if num > n {
			break
		}
		maxPos++
	}
	prev := !getBit(uint32(n), 0)
	for i := 0; i < maxPos; i++ {
		curBit := getBit(uint32(n), uint8(i))
		if curBit != prev {
			prev = !prev
			continue
		}
		return false
	}
	return true
}

func getBit(num uint32, pos uint8) bool {
	if num&(1<<pos) > 0 {
		return true
	}
	return false
}
