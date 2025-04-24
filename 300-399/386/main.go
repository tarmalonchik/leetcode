package main

func lexicalOrder(n int) []int {
	out := make([]int, 0, n)
	current := 1
	for range n {
		out = append(out, current)
		current = next(current, n, true)
	}
	return out
}

func next(current, maxNum int, up bool) int {
	if up && current*10 <= maxNum {
		return current * 10
	}

	if current%10 != 9 {
		if current+1 > maxNum {
			return next(current/10, maxNum, false)
		}
		return current + 1
	}

	if current < 10 {
		return -1
	}
	return next(current/10, maxNum, false)
}
