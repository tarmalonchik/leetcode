package main

func minCostClimbingStairs(cost []int) int {
	distances := make([]int, len(cost))

	for i := len(cost) - 1; i >= 0; i-- {
		if i == len(cost)-1 || i == len(cost)-2 {
			distances[i] = cost[i]
			continue
		}
		distances[i] = cost[i] + minVal(distances[i+1], distances[i+2])
	}
	return minVal(distances[0], distances[1])
}

func minVal(a, b int) int {
	if a < b {
		return a
	}
	return b
}
