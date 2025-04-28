package main

import (
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	out := 0
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	curPos := 0
	for {
		if curPos+1 >= len(intervals) {
			break
		}
		if intervals[curPos+1][0] >= intervals[curPos][1] && intervals[curPos+1][1] >= intervals[curPos][1] {
			curPos++
			continue
		}
		if ok, val := isInside(intervals[curPos], intervals[curPos+1]); ok {
			intervals[curPos+1] = val
			out++
			curPos++
			continue
		}
		if intervals[curPos][1] >= intervals[curPos+1][1] {
			out++
			curPos++
			continue
		}
		intervals[curPos+1] = intervals[curPos]
		out++
		curPos++
	}
	return out
}

func isInside(a, b []int) (bool, []int) {
	if a[0] >= b[0] && a[1] <= b[1] {
		return true, a
	}
	if b[0] >= a[0] && b[1] <= a[1] {
		return true, b
	}
	return false, nil
}
