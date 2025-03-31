package main

func reverseWords(s string) string {
	pos := 0
	runes := []rune(s)

	for {
		pos1, pos2 := getNextWord(runes, pos)
		if pos1 == -1 {
			break
		}
		reverseWord(runes, pos1, pos2)
		pos = pos2 + 1
	}
	return string(runes)
}

func reverseWord(s []rune, start, end int) {
	for {
		if start >= end {
			break
		}
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}

func getNextWord(s []rune, start int) (pos1, pos2 int) {
	if start > len(s)-1 {
		return -1, -1
	}

	// skip leading spaces
	for i := start; i < len(s); i++ {
		if string(s[i]) != " " {
			if i == len(s)-1 {
				return -1, -1
			}
			pos1 = i
			break
		}
	}

	for i := pos1; i < len(s); i++ {
		if string(s[i]) == " " {
			pos2 = i - 1
			break
		}
		if i == len(s)-1 {
			pos2 = len(s) - 1
			break
		}
	}
	return pos1, pos2
}
