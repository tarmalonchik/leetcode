package main

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 || matrix[0][0] > target || matrix[len(matrix)-1][len(matrix[0])-1] < target {
		return false
	}

	lMax := findLeftMax(matrix, target)
	rMin := findRightMin(matrix[:lMax+1], target)
	upMax := findUpMax(matrix, target)
	downMin := findDownMin(matrix, target)

	for i := rMin; i <= lMax; i++ {
		if findNum(matrix[i][downMin:upMax+1], target) {
			return true
		}
	}
	return false
}

func findNum(input []int, target int) bool {
	if len(input) == 0 {
		return false
	}
	minNum := 0
	maxNum := len(input) - 1
	for {
		if maxNum-minNum <= 1 {
			if input[minNum] == target {
				return true
			}
			if input[maxNum] == target {
				return true
			}
			return false
		}
		center := minNum + ((maxNum - minNum) / 2)
		if input[center] >= target {
			maxNum = center
		} else {
			minNum = center
		}
	}
}

func findDownMin(matrix [][]int, target int) (min int) {
	minNum := 0
	maxNum := len(matrix[0]) - 1
	lastRow := len(matrix) - 1
	for {
		if maxNum-minNum <= 1 {
			if matrix[lastRow][minNum] >= target {
				return minNum
			}
			if matrix[lastRow][maxNum] >= target {
				return maxNum
			}
			return minNum
		}
		center := minNum + ((maxNum - minNum) / 2)
		if matrix[lastRow][center] <= target {
			minNum = center
		} else {
			maxNum = center
		}
	}
}

func findUpMax(matrix [][]int, target int) (max int) {
	minNum := 0
	maxNum := len(matrix[0]) - 1
	for {
		if maxNum-minNum <= 1 {
			if matrix[0][maxNum] <= target {
				return maxNum
			}
			if matrix[0][minNum] <= target {
				return minNum
			}
			return maxNum
		}
		center := minNum + ((maxNum - minNum) / 2)
		if matrix[0][center] <= target {
			minNum = center
		} else {
			maxNum = center
		}
	}
}

func findLeftMax(matrix [][]int, target int) (max int) {
	minNum := 0
	maxNum := len(matrix) - 1
	for {
		if maxNum-minNum <= 1 {
			if matrix[maxNum][0] <= target {
				return maxNum
			}
			if matrix[minNum][0] <= target {
				return minNum
			}
			return maxNum
		}
		center := minNum + ((maxNum - minNum) / 2)
		if matrix[center][0] > target {
			maxNum = center
		} else {
			minNum = center
		}
	}
}

func findRightMin(matrix [][]int, target int) (min int) {
	if matrix[0][len(matrix[0])-1] >= target {
		return 0
	}

	lastColumn := len(matrix[0]) - 1

	minNum := 0
	maxNum := len(matrix[0]) - 1

	for {
		if maxNum-minNum <= 1 {
			if matrix[minNum][lastColumn] >= target {
				return minNum
			}
			if matrix[maxNum][lastColumn] >= target {
				return maxNum
			}
			return minNum
		}
		center := minNum + ((maxNum - minNum) / 2)
		if matrix[center][lastColumn] >= target {
			maxNum = center
		} else {
			minNum = center
		}
	}
}
