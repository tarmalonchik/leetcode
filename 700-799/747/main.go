package main

func dominantIndex(nums []int) int {
	m := maxNums{}
	for i := range nums {
		m.add(nums[i], i)
	}
	return m.getIdx()
}

type maxNums struct {
	setNum uint8
	maxOne int
	maxTwo int
	maxIDX int
}

func (m *maxNums) getIdx() int {
	if m.maxOne >= m.maxTwo*2 {
		return m.maxIDX
	}
	return -1
}

func (m *maxNums) add(num, idx int) {
	if m.setNum == 0 {
		m.maxOne = num
		m.maxIDX = idx
		m.setNum++
		return
	}
	if m.setNum == 1 {
		if num >= m.maxOne {
			m.maxTwo = m.maxOne
			m.maxOne = num
			m.maxIDX = idx
			m.setNum++
			return
		}
		m.maxTwo = num
		m.setNum++
	}
	if num > m.maxOne {
		m.maxTwo = m.maxOne
		m.maxOne = num
		m.maxIDX = idx
		return
	}
	if num > m.maxTwo {
		m.maxTwo = num
		return
	}
}
