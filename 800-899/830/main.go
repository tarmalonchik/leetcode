package main

func largeGroupPositions(s string) [][]int {
	if len(s) < 3 {
		return nil
	}
	out := make([][]int, 0)
	prev := s[0]
	groupStart := 0

	for i := 1; i < len(s); i++ {
		if prev == s[i] {
			continue
		}
		if i-groupStart >= 3 {
			out = append(out, []int{groupStart, i - 1})
		}
		prev = s[i]
		groupStart = i
	}
	if len(s)-groupStart >= 3 {
		out = append(out, []int{groupStart, len(s) - 1})
	}
	return out
}
