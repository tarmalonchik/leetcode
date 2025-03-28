package main

func findPoisonedDuration(timeSeries []int, duration int) int {
	if len(timeSeries) == 0 {
		return 0
	}

	total := 0
	idx := 1
	lastVal := timeSeries[0]

	for {
		if idx > len(timeSeries)-1 {
			break
		}
		total += minNum(duration, diff(lastVal, timeSeries[idx]))
		lastVal = timeSeries[idx]
		idx++
	}
	return total + duration
}

func minNum(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
