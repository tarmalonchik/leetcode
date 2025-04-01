package main

func maxCount(m int, n int, ops [][]int) int {
	x, y := findMinXY(ops)
	if x == -1 {
		x = m
	}
	if y == -1 {
		y = n
	}
	return x * y
}

func findMinXY(ops [][]int) (int, int) {
	minX := -1
	minY := -1
	for i := range ops {
		if ops[i][0] != 0 {
			if minX == -1 {
				minX = ops[i][0]
			} else if minX > ops[i][0] {
				minX = ops[i][0]
			}
		}
		if ops[i][1] != 0 {
			if minY == -1 {
				minY = ops[i][1]
			} else if minY > ops[i][1] {
				minY = ops[i][1]
			}
		}
	}
	return minX, minY
}
