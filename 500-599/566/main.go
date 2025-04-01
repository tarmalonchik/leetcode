package main

func matrixReshape(mat [][]int, r int, c int) [][]int {
	if len(mat) == 0 || len(mat[0]) == 0 {
		return mat
	}
	if len(mat)*len(mat[0]) != r*c {
		return mat
	}

	newMat := make([][]int, 0, r)
	for i := 0; i < r*c; i++ {
		newMat = fillNext(newMat, c, getNext(mat, i))
	}
	return newMat
}

func fillNext(mat [][]int, rowSize, num int) [][]int {
	if len(mat) == 0 || len(mat[len(mat)-1]) >= rowSize {
		mat = append(mat, make([]int, 0, rowSize))
	}
	mat[len(mat)-1] = append(mat[len(mat)-1], num)
	return mat
}

func getNext(mat [][]int, i int) int {
	col := i % len(mat[0])
	row := i / len(mat[0])
	return mat[row][col]
}
