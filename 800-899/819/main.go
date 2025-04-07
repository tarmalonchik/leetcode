package main

import (
	"strings"
)

func mostCommonWord(paragraph string, banned []string) string {
	ban := make(map[string]interface{})
	words := make(map[string]int)
	for i := range banned {
		ban[banned[i]] = nil
	}
	paragraphByte := []byte(paragraph)

	out := ""
	maxNum := 0
	pos1 := 0
	pos2 := 0

	for {
		pos1, pos2 = nextWord(paragraphByte, pos1)
		if pos1 == -1 {
			break
		}
		lowerVal := strings.ToLower(string(paragraphByte[pos1:pos2]))
		pos1 = pos2
		if _, ok := ban[lowerVal]; ok {
			continue
		}
		words[lowerVal]++
		count := words[lowerVal]
		if count > maxNum {
			maxNum = count
			out = lowerVal
		}
	}
	return out
}

func nextWord(in []byte, pos int) (pos1, pos2 int) {
	for {
		if pos > len(in)-1 {
			return -1, -1
		}
		if !isSymbol(in[pos]) {
			pos++
			continue
		}
		break
	}
	pos1 = pos
	pos2 = pos
	for {
		if pos2 > len(in)-1 {
			return pos1, pos2
		}
		if !isSymbol(in[pos2]) {
			break
		}
		pos2++
	}
	return pos1, pos2
}

func isSymbol(in byte) bool {
	if in >= 97 && in <= 122 {
		return true
	}
	if in >= 65 && in <= 90 {
		return true
	}
	return false
}
