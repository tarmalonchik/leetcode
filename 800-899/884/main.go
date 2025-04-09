package main

import (
	"strings"
)

func uncommonFromSentences(s1 string, s2 string) []string {
	items := strings.Split(s1+" "+s2, " ")
	mpSingle := make(map[string]interface{})
	mpMultiple := make(map[string]interface{})

	for i := range items {
		if _, ok := mpMultiple[items[i]]; ok {
			continue
		}
		if _, ok := mpSingle[items[i]]; ok {
			delete(mpSingle, items[i])
			mpMultiple[items[i]] = items[i]
			continue
		}
		mpSingle[items[i]] = nil
	}
	out := make([]string, 0, len(mpSingle))
	for key := range mpSingle {
		out = append(out, key)
	}
	return out
}
