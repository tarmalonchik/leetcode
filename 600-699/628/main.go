package main

func maximumProduct(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	if len(nums) == 3 {
		return nums[0] * nums[1] * nums[2]
	}

	min2 := minTwo{}
	max3 := maxThree{}

	for i := range nums {
		min2.add(nums[i])
		max3.add(nums[i])
	}

	fCase := min2.multiple() * max3.getMax()
	sCase := max3.multiple()
	if fCase > sCase {
		return fCase
	}
	return sCase
}

type minTwo struct {
	valid int
	a     int
	b     int
}

func (m *minTwo) multiple() int {
	return m.a * m.b
}

func (m *minTwo) add(num int) {
	if m.valid == 0 {
		m.valid++
		m.a = num
		return
	} else if m.valid == 1 {
		m.valid++
		m.b = num
		return
	}
	if num > m.a && num > m.b {
		return
	}
	if num < m.a && num > m.b {
		m.a = num
		return
	}
	if num < m.b && num > m.a {
		m.b = num
	}

	if m.a < m.b {
		m.b = num
	} else {
		m.a = num
	}
}

type maxThree struct {
	valid  int
	min    int
	center int
	max    int
}

func (m *maxThree) multiple() int {
	return m.min * m.center * m.max
}

func (m *maxThree) getMax() int {
	return m.max
}

func (m *maxThree) add(num int) {
	if m.valid == 0 {
		m.valid++
		m.max = num
		return
	}
	if m.valid == 1 {
		m.valid++
		if num > m.max {
			m.center = m.max
			m.max = num
			return
		}
		m.center = num
		return
	}
	if m.valid == 2 {
		m.valid++
		if num > m.max {
			m.min = m.center
			m.center = m.max
			m.max = num
			return
		}
		if num > m.center {
			m.min = m.center
			m.center = num
			return
		}
		m.min = num
		return
	}
	if num > m.max {
		m.min = m.center
		m.center = m.max
		m.max = num
		return
	}
	if num > m.center {
		m.min = m.center
		m.center = num
		return
	}
	if num > m.min {
		m.min = num
	}
}
