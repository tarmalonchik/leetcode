package main

import (
	"math/rand/v2"
)

type Solution struct {
	mp map[int][]int
}

func Constructor(nums []int) Solution {
	s := Solution{
		mp: make(map[int][]int),
	}
	for i := range nums {
		s.mp[nums[i]] = append(s.mp[nums[i]], i)
	}
	return s
}

func (this *Solution) Pick(target int) int {
	return this.mp[target][randRange(0, len(this.mp[target]))]
}

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}
