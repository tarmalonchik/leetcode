package main

func transpose(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return matrix
	}
	out := make([][]int, len(matrix[0]))
	for i := range out {
		out[i] = make([]int, len(matrix))
	}

	for y := range matrix {
		for x := range matrix[y] {
			out[x][y] = matrix[y][x]
		}
	}
	return out
}
