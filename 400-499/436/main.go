package main

import (
	"sort"
)

func findRightInterval(intervals [][]int) []int {
	if len(intervals) == 0 {
		return nil
	}
	maxNum := intervals[0][0]
	positions := make(map[int]int, len(intervals))
	for i := range intervals {
		if intervals[i][0] > maxNum {
			maxNum = intervals[i][0]
		}
		positions[intervals[i][0]] = i
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	out := make([]int, len(intervals))

	for i := range intervals {
		if intervals[i][1] > maxNum {
			out[positions[intervals[i][0]]] = -1
			continue
		}
		out[positions[intervals[i][0]]] = positions[find(intervals, intervals[i][1])]
	}
	return out
}

func find(in [][]int, num int) int {
	pos1 := 0
	pos2 := len(in) - 1

	for {
		if pos2-pos1 <= 1 {
			if in[pos1][0] >= num {
				return in[pos1][0]
			}
			return in[pos2][0]
		}
		center := pos1 + (pos2-pos1)/2
		if in[center][0] < num {
			pos1 = center
		} else {
			pos2 = center
		}
	}
}
