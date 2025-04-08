package main

func surfaceArea(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	surface := 0
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] > 0 {
				surface += 2
			}
			surface += grid[y][x] * edgesCount(grid, x, y)
			if x < len(grid[y])-1 {
				surface += abs(grid[y][x], grid[y][x+1])
			}
			if y < len(grid)-1 {
				surface += abs(grid[y][x], grid[y+1][x])
			}
		}
	}
	return surface
}

func abs(a, b int) int {
	if a-b < 0 {
		return b - a
	}
	return a - b
}

func edgesCount(grid [][]int, x, y int) int {
	edges := 0
	if x == 0 {
		edges++
	}
	if y == 0 {
		edges++
	}
	if x == len(grid[0])-1 {
		edges++
	}
	if y == len(grid)-1 {
		edges++
	}
	return edges
}
