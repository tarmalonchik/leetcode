package main

func reverseStr(s string, k int) string {
	sByte := []byte(s)
	pos := 0
	for {
		if pos >= len(sByte) {
			break
		}
		if pos+k >= len(sByte) {
			reverseRange(sByte, pos, len(sByte))
			break
		}
		reverseRange(sByte, pos, pos+k)
		pos += 2 * k
	}
	return string(sByte)
}

func reverseRange(s []byte, pos1, pos2 int) {
	pos2--
	for {
		if pos2 <= pos1 {
			break
		}
		s[pos1], s[pos2] = s[pos2], s[pos1]
		pos1++
		pos2--
	}
}
