package main

import (
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	var out int

	gPos := len(g) - 1
	sPos := len(s) - 1
	for {
		if gPos < 0 || sPos < 0 {
			break
		}
		if s[sPos] >= g[gPos] {
			out++
			sPos--
			gPos--
			continue
		}
		gPos--
	}
	return out
}
