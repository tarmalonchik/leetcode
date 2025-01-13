package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum2([]int{1, 1}, 1))
}

func combinationSum2(candidates []int, target int) [][]int {
	out := make(output, 0)
	manager := newCandidateManager(candidates)
	out.recursion(target, manager, []int{})
	return out
}

func (o *output) recursion(target int, manager candidateManager, out []int) {
	if target < 0 {
		return
	}
	if target == 0 {
		localOut := make([]int, len(out))
		copy(localOut, out)
		*o = append(*o, localOut)
		return
	}

	if manager.len() == 0 {
		return
	}

	if manager.isPhantom() {
		mg := manager
		mg.idx++
		o.recursion(target-mg.getCurrent(), mg, append(out, mg.getCurrent()))
	} else {
		mg := manager
		current := mg.getCurrent()
		mg = mg.goNext()
		o.recursion(target-current, mg, append(out, current))
	}
	manager = manager.goNext()
	o.recursion(target, manager, out)
}

type output [][]int

type candidateManager struct {
	candidates []int
	count      int
	length     int
	idx        int
}

func (c *candidateManager) getCurrent() int {
	return c.candidates[c.idx]
}

func (c *candidateManager) isPhantom() bool {
	if c.length == 0 {
		panic("some")
	}
	if len(c.candidates)-1 == c.idx {
		return false
	}
	if c.candidates[c.idx] == c.candidates[c.idx+1] {
		return true
	}
	return false
}

func (c *candidateManager) goNext() candidateManager {
	if c.length == 0 {
		return *c
	}

	if c.length == 1 {
		c.length = 0
		return *c
	}

	cur := c.candidates[c.idx]
	for i := c.idx + 1; i < len(c.candidates); i++ {
		if c.candidates[i] != cur {
			c.idx = i
			c.length--
			return *c
		}
		if i == len(c.candidates)-1 {
			c.length = 0
			return *c
		}
	}
	panic("invalid behaviour")
}

func (c *candidateManager) len() int {
	return c.length
}

func (c *candidateManager) countMe() {
	if len(c.candidates) == 0 {
		return
	}
	cur := c.candidates[0]
	count := 1

	for i := 1; i < len(c.candidates); i++ {
		if cur != c.candidates[i] {
			count++
			cur = c.candidates[i]
			continue
		}
	}
	c.count = count
}

func newCandidateManager(candidates []int) candidateManager {
	sort.Ints(candidates)
	c := candidateManager{
		candidates: candidates,
	}
	c.countMe()
	c.length = c.count
	return c
}
