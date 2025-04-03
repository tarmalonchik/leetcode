package main

import (
	"sort"
)

type KthLargest struct {
	arr []int
}

func Constructor(k int, nums []int) KthLargest {
	item := KthLargest{
		arr: make([]int, 0, k),
	}
	for i := range nums {
		item.Add(nums[i])
	}
	return item
}

func (this *KthLargest) Add(val int) int {
	if len(this.arr) < cap(this.arr) {
		this.arr = append(this.arr, val)
		sort.Ints(this.arr)
		return this.arr[0]
	}

	if this.arr[0] > val {
		return this.arr[0]
	}
	this.arr[0] = val
	sort.Ints(this.arr)
	return this.arr[0]
}
