package main

func toHex(num int) string {
	if num < 0 {
		return toHexPositive(4294967296 + num)
	}
	return toHexPositive(num)
}

func toHexPositive(num int) string {
	if num < 0 {
		panic("unexpected")
	}
	preOut := make([]byte, 0)

	for {
		part := num % 16
		num = num / 16
		preOut = append(preOut, singleNumToHex(part))
		if num == 0 {
			break
		}
	}

	start := 0
	end := len(preOut) - 1
	for {
		if start >= end {
			break
		}
		preOut[start], preOut[end] = preOut[end], preOut[start]
		start++
		end--
	}

	return string(preOut)
}

func singleNumToHex(num int) byte {
	if num < 0 || num > 15 {
		panic("unexpected")
	}
	if num <= 9 {
		return 48 + byte(num)
	}
	return byte(num) + 87
}
