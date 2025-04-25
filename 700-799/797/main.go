package main

func allPathsSourceTarget(graph [][]int) [][]int {
	in := []int{0}
	return recursion(graph, in, 0, len(graph)-1)
}

func recursion(graph [][]int, in []int, offset, target int) [][]int {
	var out [][]int

	if offset > len(graph)-1 {
		return nil
	}

	for i := range graph[offset] {
		newIn := make([]int, len(in), len(in)+1)
		copy(newIn, in)
		val := graph[offset][i]
		newIn = append(newIn, val)
		if val == target {
			out = append(out, newIn)
		} else {
			out = append(out, recursion(graph, newIn, val, target)...)
		}
	}
	return out
}
