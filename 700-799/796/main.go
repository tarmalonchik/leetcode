package main

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	if len(s) == 0 || len(s) == 1 {
		if len(s) == 1 {
			return s == goal
		}
		return true
	}
	sSum := 0
	goalSum := 0

	for i := range s {
		sSum += int(s[i])
	}
	for i := range goal {
		goalSum += int(goal[i])
	}

	if sSum != goalSum {
		return false
	}

	for i := 0; i < len(s); i++ {
		for j := range s {
			if s[j] != getCharOffset(goal, j+i) {
				break
			}
			if j == len(s)-1 {
				return true
			}
		}
	}
	return false
}

func getCharOffset(in string, offset int) byte {
	return in[offset%len(in)]
}
