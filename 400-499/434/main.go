package main

func countSegments(s string) int {
	segments := 0
	var lastOne = byte(' ')
	for _, val := range []byte(s) {
		if val == ' ' && lastOne != ' ' {
			segments++
			lastOne = ' '
		}
		lastOne = val
	}
	if lastOne != ' ' {
		segments++
	}
	return segments
}
