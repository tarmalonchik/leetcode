package main

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	sqrt := computeSqr(ax1, ax2, ay1, ay2) + computeSqr(bx1, bx2, by1, by2)

	xNum := returnIntersectionSize(ax1, ax2, bx1, bx2)
	if xNum == 0 {
		return sqrt
	}
	yNum := returnIntersectionSize(ay1, ay2, by1, by2)
	if yNum == 0 {
		return sqrt
	}
	return sqrt - (xNum * yNum)
}

func computeSqr(x1, x2, y1, y2 int) (value int) {
	return (x2 - x1) * (y2 - y1)
}

func returnIntersectionSize(a1, a2, b1, b2 int) (value int) {
	if a2 <= b1 || b2 <= a1 { // no intersection
		return 0
	}
	if (a1 >= b1 && a2 <= b2) || (b1 >= a1 && b2 <= a2) { // one inside other
		return minNum(a2-a1, b2-b1)
	}
	if (a1 >= b1 && a1 <= b2) || (b1 >= a1 && b1 <= a2) { // intersection
		if a1 >= b1 && a1 <= b2 {
			return b2 - a1
		} else {
			return a2 - b1
		}
	}
	panic("invalid case")
}

func minNum(a, b int) int {
	if a < b {
		return a
	}
	return b
}
