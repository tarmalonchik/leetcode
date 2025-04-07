package main

func backspaceCompare(s string, t string) bool {
	pos1 := len(s) - 1
	pos2 := len(t) - 1

	last := pos1 + pos2 + 1
	for {
		if last == pos1+pos2 {
			break
		}
		last = pos1 + pos2
		pos1 = skip(s, pos1)
		pos2 = skip(t, pos2)
		if pos1 == -1 || pos2 == -1 {
			if pos1 == -1 && pos2 == -1 {
				return true
			}
			return false
		}
		if s[pos1] == t[pos2] {
			pos1--
			pos2--
			continue
		}
		return false
	}
	return false
}

func skip(in string, pos int) int {
	skipCount := 0
	for i := pos; i >= 0; i-- {
		if in[i] == '#' {
			skipCount++
			continue
		}
		if skipCount > 0 {
			skipCount--
			continue
		}
		return i
	}
	return -1
}
