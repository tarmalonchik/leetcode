package main

func reverseOnlyLetters(s string) string {
	sByte := []byte(s)

	pos1 := 0
	pos2 := len(sByte) - 1

	for {
		if pos1 >= pos2 {
			break
		}
		if !isLetter(sByte[pos1]) {
			pos1++
			continue
		}
		if !isLetter(sByte[pos2]) {
			pos2--
			continue
		}
		sByte[pos1], sByte[pos2] = sByte[pos2], sByte[pos1]
		pos1++
		pos2--
	}
	return string(sByte)
}

func isLetter(in byte) bool {
	if in >= 97 && in <= 122 {
		return true
	}
	if in >= 65 && in <= 90 {
		return true
	}
	return false
}
