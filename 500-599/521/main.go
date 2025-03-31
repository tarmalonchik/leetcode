package main

func findLUSlength(a string, b string) int {
	if a != b {
		return getMax(len(a), len(b))
	}
	return -1
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
