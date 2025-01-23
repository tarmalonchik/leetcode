package main

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	out := make(output, 0)
	sort.Ints(candidates)
	out.recursion(target, candidates, []int{})
	return out
}

type output [][]int

func (o *output) recursion(target int, candidates []int, out []int) {
	if len(candidates) == 0 {
		return
	}
	if target < 0 {
		return
	}
	if target == 0 {
		localOut := make([]int, len(out))
		copy(localOut, out)
		*o = append(*o, localOut)
		return
	}

	o.recursion(target-candidates[0], candidates, append(out, candidates[0]))

	if len(candidates) > 1 {
		o.recursion(target-candidates[1], candidates[1:], append(out, candidates[1]))
	}
	if len(candidates) > 2 {
		o.recursion(target, candidates[2:], out)
	}
}
