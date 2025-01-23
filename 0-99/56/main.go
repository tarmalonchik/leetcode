package main

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	out := make([][]int, 0)

	for i := range intervals {
		if len(out) == 0 || out[len(out)-1][1] < intervals[i][0] {
			out = append(out, intervals[i])
			continue
		}
		if out[len(out)-1][0] <= intervals[i][0] && out[len(out)-1][1] >= intervals[i][0] {
			out[len(out)-1][1] = getMax(out[len(out)-1][1], intervals[i][1])
		}
	}
	return out
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
