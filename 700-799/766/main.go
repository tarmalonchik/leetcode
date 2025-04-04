package main

func isToeplitzMatrix(matrix [][]int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return true
	}
	x := 0
	y := len(matrix) - 1
	return check(matrix, []pos{{x: x, y: y}})
}

func check(matrix [][]int, inPos []pos) bool {
	set := false
	val := 0
	newPositions := make([]pos, 0)
	for i := range inPos {
		if !set {
			val1, ok := exists(matrix, inPos[i].y-1, inPos[i].x)
			if ok {
				newPositions = append(newPositions, pos{y: inPos[i].y - 1, x: inPos[i].x})
				set = true
				val = val1
			}
		}
		val2, ok := exists(matrix, inPos[i].y, inPos[i].x+1)
		if ok {
			newPositions = append(newPositions, pos{y: inPos[i].y, x: inPos[i].x + 1})
			if !set {
				set = true
				val = val2
			} else if val != val2 {
				return false
			}
		}
	}
	if len(newPositions) == 0 {
		return true
	}
	return check(matrix, newPositions)
}

func exists(matrix [][]int, y, x int) (int, bool) {
	if x < 0 || y < 0 || x > len(matrix[0])-1 || y > len(matrix)-1 {
		return 0, false
	}
	return matrix[y][x], true
}

type pos struct {
	y, x int
}
