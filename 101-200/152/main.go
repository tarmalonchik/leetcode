package main

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxNum := nums[0]
	prevCont := container{
		case1: nums[0],
		case2: nums[0],
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] == 0 {
			if maxNum < 0 {
				maxNum = 0
			}
			prevCont = container{}
			continue
		}
		currentCont := prevCont
		currentCont.addNum(nums[i])
		if currentCont.getMax() > maxNum {
			maxNum = currentCont.getMax()
		}
		prevCont = currentCont
	}
	return maxNum
}

type container struct {
	case1 int
	case2 int
}

func (c *container) isZero() bool {
	if c.case1 == 0 && c.case2 == 0 {
		return true
	}
	return false
}

func (c *container) addNum(in int) {
	num1 := in * c.case1
	num2 := in * c.case2

	c.case1 = in
	c.case2 = in

	if c.case1 < num1 {
		c.case1 = num1
	}
	if c.case1 < num2 {
		c.case1 = num2
	}

	if c.case2 > num1 {
		c.case2 = num1
	}
	if c.case2 > num2 {
		c.case2 = num2
	}
}

func (c *container) getMax() int {
	if c.case1 > c.case2 {
		return c.case1
	}
	return c.case2
}
