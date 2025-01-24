package main

func reverseString(s []byte) {
	if len(s) == 0 {
		return
	}
	pos1 := 0
	pos2 := len(s) - 1
	for {
		if pos1 >= pos2 {
			break
		}
		s[pos1], s[pos2] = s[pos2], s[pos1]
		pos1++
		pos2--
	}
}
