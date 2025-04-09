package main

func projectionArea(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	sum := 0

	for y := range grid {
		maxX := grid[y][0]
		for x := range grid[y] {
			if grid[y][x] > 0 {
				sum++
			}
			if grid[y][x] > maxX {
				maxX = grid[y][x]
			}
		}
		sum += maxX
	}

	for x := range grid[0] {
		maxY := grid[0][x]
		for y := range grid {
			if grid[y][x] > maxY {
				maxY = grid[y][x]
			}
		}
		sum += maxY
	}
	return sum
}
