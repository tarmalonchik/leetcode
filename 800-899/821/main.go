package main

func shortestToChar(s string, c byte) []int {
	distance := -1
	out := make([]int, len(s))
	for i := range s {
		if s[i] == c {
			distance = 0
			continue
		}
		if distance != -1 {
			distance++
		}
		out[i] = distance
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == c {
			distance = 0
			continue
		}
		if distance != -1 {
			distance++
		}
		if distance < out[i] || out[i] == -1 {
			out[i] = distance
		}
	}
	return out
}
