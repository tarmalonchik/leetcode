package main

import (
	"strings"
)

func findWords(words []string) []string {
	first := "qwertyuiop"
	second := "asdfghjkl"
	third := "zxcvbnm"

	mp := make(map[string]uint8, len(first)+len(second)+len(third))
	fill(first, second, third, mp)

	var out []string
	for i := range words {
		if check(mp, words[i]) {
			out = append(out, words[i])
		}
	}
	return out
}

func fill(first, second, third string, mp map[string]uint8) {
	for i := range first {
		mp[string(first[i])] = 1
	}
	for i := range second {
		mp[string(second[i])] = 2
	}
	for i := range third {
		mp[string(third[i])] = 3
	}
}

func check(mp map[string]uint8, in string) bool {
	last := uint8(0)
	for i := range in {
		if last == 0 {
			last = mp[strings.ToLower(string(in[i]))]
			continue
		}
		if last != mp[strings.ToLower(string(in[i]))] {
			return false
		}
	}
	return true
}
