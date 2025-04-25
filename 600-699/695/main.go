package main

func maxAreaOfIsland(grid [][]int) int {
	mp := make(map[position]interface{})
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == 1 {
				mp[position{x: x, y: y}] = nil
			}
		}
	}

	maxOut := 0

	for {
		if len(mp) == 0 {
			break
		}
		for key := range mp {
			count := process(key, mp)
			if count > maxOut {
				maxOut = count
			}
			break
		}
	}
	return maxOut
}

func process(pos position, mp map[position]interface{}) int {
	out := 0

	if pos.x < 0 || pos.y < 0 {
		return 0
	}
	if _, ok := mp[pos]; ok {
		out += 1
	} else {
		return 0
	}
	delete(mp, pos)
	out += process(position{pos.x, pos.y - 1}, mp)
	out += process(position{pos.x, pos.y + 1}, mp)
	out += process(position{pos.x - 1, pos.y}, mp)
	out += process(position{pos.x + 1, pos.y}, mp)
	return out
}

type position struct {
	x, y int
}
