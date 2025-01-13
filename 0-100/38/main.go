package main

import (
	"strconv"
)

func countAndSay(n int) string {
	start := "1"

	for i := 0; i < n-1; i++ {
		start = oneOperation(start)
	}
	return start
}

func oneOperation(in string) (out string) {
	char := in[0]
	count := 1
	for i := 1; i < len(in); i++ {
		if in[i] == char {
			count++
			continue
		}
		out += strconv.Itoa(count) + string(char)
		count = 1
		char = in[i]
	}
	return out + strconv.Itoa(count) + string(char)
}
