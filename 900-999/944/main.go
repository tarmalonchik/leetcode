package main

func minDeletionSize(strs []string) int {
	if len(strs) < 2 || len(strs[0]) == 0 {
		return 0
	}
	var out int

	for x := range strs[0] {
		for y := 1; y <= len(strs)-1; y++ {
			if strs[y][x] < strs[y-1][x] {
				out++
				break
			}
		}
	}
	return out
}
