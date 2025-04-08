package main

func numberOfLines(widths []int, s string) []int {
	if len(s) == 0 {
		return []int{1, 0}
	}
	lines := 1
	pixels := 100

	for i := range s {
		if pixels-getWidth(s[i], widths) >= 0 {
			pixels -= getWidth(s[i], widths)
		} else {
			lines++
			pixels = 100 - getWidth(s[i], widths)
		}
	}
	return []int{lines, 100 - pixels}
}

func getWidth(in byte, widths []int) int {
	return widths[int(in-97)]
}
