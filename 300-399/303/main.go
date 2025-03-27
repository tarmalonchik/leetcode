package main

type NumArray struct {
	nums []int
	sums []int
}

func Constructor(nums []int) NumArray {
	out := NumArray{
		nums: nums,
	}
	sum := 0

	for i := range nums {
		sum += nums[i]
		out.sums = append(out.sums, sum)
	}
	return out
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.sums[right] - this.sums[left] + this.nums[left]
}
