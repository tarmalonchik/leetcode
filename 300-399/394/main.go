package main

import (
	"fmt"
)

func main() {
	fmt.Println(decodeString("3[a2[c]]"))
}

func decodeString(s string) string {
	out := make([]byte, 0)
	out, _ = recursive(s, out, 0)
	return string(out)
}

func recursive(s string, out []byte, skip int) ([]byte, int) {
	num := 0

	for i := 0; i < len(s); i++ {
		dig := toDig(s[i])
		if dig != -1 {
			num *= 10
			num += dig
			continue
		}
		if s[i] == '[' {
			skip++
			continue
		}
		if s[i] == ']' {
			if skip == 0 {
				return out, i
			}
			skip--
			continue
		}

		startPos := 0
		for range num {
			startPos = i
			for {
				if toDig(s[startPos]) != -1 {
					offset := 0
					out, offset = recursive(s[startPos:], out, 0)
					if offset == -1 {
						break
					}
					startPos += offset
				}
				if s[startPos] == ']' {
					break
				}
				out = append(out, s[startPos])
				startPos++
			}
		}
		i = startPos
		num = 0
	}
	return out, -1
}

func toDig(in byte) int {
	if in >= 48 && in <= 57 {
		return int(in) - 48
	}
	return -1
}
