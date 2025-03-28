package main

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	addNegative := false
	if num < 0 {
		num = -num
		addNegative = true
	}
	out := make([]byte, 0)

	for {
		if num == 0 {
			break
		}
		a := num % 7
		out = append(out, numToByte(a))
		num = num / 7
	}

	if addNegative {
		out = append(out, '-')
	}

	pos1 := 0
	pos2 := len(out) - 1

	for {
		if pos2 <= pos1 {
			break
		}
		out[pos1], out[pos2] = out[pos2], out[pos1]
		pos1++
		pos2--
	}
	return string(out)
}

func numToByte(in int) byte {
	return 48 + byte(in)
}
