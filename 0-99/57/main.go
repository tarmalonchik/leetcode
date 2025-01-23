package main

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	idx := findIdx(intervals, newInterval)
	if idx == len(intervals)-1 {
		if overlap(newInterval, intervals[idx]) {
			intervals[idx] = merge(intervals[idx], newInterval)
			return intervals
		}
		if newInterval[0] > intervals[idx][1] {
			return append(intervals, newInterval)
		} else {
			out := append(intervals, []int{0, 0})
			out[len(out)-1][0] = out[idx][0]
			out[len(out)-1][1] = out[idx][1]
			out[idx][0] = newInterval[0]
			out[idx][1] = newInterval[1]
			return out
		}
	}
	if idx == 0 && newInterval[1] < intervals[idx][0] {
		return append([][]int{newInterval}, intervals...)
	}

	newIdx := idx
	for {
		if newIdx == len(intervals) {
			break
		}
		if overlap(intervals[newIdx], newInterval) {
			newInterval = merge(intervals[newIdx], newInterval)
			newIdx++
			continue
		}
		break
	}

	part2 := make([][]int, len(intervals[newIdx:]))
	for i := range intervals[newIdx:] {
		part2[i] = make([]int, 2)
		part2[i][0] = intervals[newIdx:][i][0]
		part2[i][1] = intervals[newIdx:][i][1]
	}

	out := append(intervals[:idx], newInterval)
	out = append(out, part2...)
	return out
}

func overlap(a, b []int) bool {
	if a[0] <= b[1] && a[0] >= b[0] {
		return true
	}
	if a[1] <= b[1] && a[1] >= b[0] {
		return true
	}
	if b[0] <= a[1] && b[0] >= a[0] {
		return true
	}
	if b[1] <= a[1] && b[1] >= a[0] {
		return true
	}
	return false
}

func merge(a, b []int) []int {
	out := make([]int, 2)
	out[0] = getMin(a[0], b[0])
	out[1] = getMax(a[1], b[1])
	return out
}

func findIdx(intervals [][]int, newInterval []int) int {
	minVal := 0
	maxVal := len(intervals) - 1

	for {
		center := minVal + (maxVal-minVal)/2
		if maxVal-minVal <= 1 {
			if overlap(intervals[minVal], newInterval) {
				return minVal
			}
			if overlap(intervals[maxVal], newInterval) {
				return maxVal
			}
			if newInterval[1] < intervals[minVal][0] {
				return minVal
			}
			return maxVal
		}
		if overlap(intervals[center], newInterval) {
			maxVal = center
			continue
		}
		if intervals[center][0] > newInterval[1] {
			maxVal = center
		} else {
			minVal = center
		}
	}
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
