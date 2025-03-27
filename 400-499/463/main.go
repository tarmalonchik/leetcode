package main

func islandPerimeter(grid [][]int) int {
	out := 0
	for i := range grid {
		for j := range grid[i] {
			out += getSingle(grid, i, j)
		}
	}
	return out
}

func getSingle(grid [][]int, i, j int) int {
	var out int

	if grid[i][j] == 0 {
		return 0
	}
	if i == 0 || grid[i-1][j] == 0 {
		out++
	}
	if i == len(grid)-1 || grid[i+1][j] == 0 {
		out++
	}
	if j == 0 || grid[i][j-1] == 0 {
		out++
	}
	if j == len(grid[0])-1 || grid[i][j+1] == 0 {
		out++
	}
	return out
}
