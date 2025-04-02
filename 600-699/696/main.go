package main

func countBinarySubstrings(s string) int {
	if len(s) <= 1 {
		return 0
	}

	groups := make([]int, 0)
	groups = append(groups, 1)
	out := 0

	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			if len(groups) >= 2 {
				out += minNum(groups[len(groups)-2], groups[len(groups)-1])
			}
			groups = append(groups, 1)
		} else {
			groups[len(groups)-1]++
		}
	}
	if len(groups) >= 2 {
		out += minNum(groups[len(groups)-2], groups[len(groups)-1])
	}
	return out
}

func minNum(a, b int) int {
	if a < b {
		return a
	}
	return b
}
