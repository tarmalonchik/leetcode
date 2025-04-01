package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}

	pos1 := 0
	pos2 := 0

	pos1, pos2 = findNextWindow(flowerbed, pos1)
	if pos1 < 0 {
		return false
	}

	for {
		n -= canPlace(pos2-pos1+1, edges(pos1, pos2, len(flowerbed)))
		if n <= 0 {
			return true
		}
		pos1, pos2 = findNextWindow(flowerbed, pos2+1)
		if pos1 < 0 {
			break
		}
	}
	return false
}

func edges(pos1, pos2, len int) (edges int) {
	if pos1 == 0 {
		edges++
	}
	if pos2 == len-1 {
		edges++
	}
	return edges
}

func canPlace(slots int, edgeNum int) int {
	if slots == 0 {
		return 0
	}
	if slots%2 == 0 {
		return canPlaceEven(slots, edgeNum)
	}
	return canPlaceOdd(slots, edgeNum)
}

func canPlaceEven(slots int, edgeNum int) int {
	if edgeNum > 0 {
		return slots/2 - 1 + 1
	}
	return slots/2 - 1
}

func canPlaceOdd(slots int, edgeNum int) int {
	if edgeNum >= 2 {
		return slots/2 + 1
	}
	return slots / 2
}

func findNextWindow(in []int, pos int) (pos1, pos2 int) {
	for {
		if pos >= len(in) || in[pos] == 0 {
			break
		}
		if in[pos] == 1 {
			pos++
			continue
		}
	}

	if pos >= len(in) {
		return -1, -1
	}

	pos1 = pos
	for {
		if pos >= len(in) {
			return pos1, len(in) - 1
		}
		if in[pos] == 1 {
			return pos1, pos - 1
		}
		pos++
	}
}
