package main

func decodeString(s string) string {
	return string(decodeByte([]byte(s)))
}

func decodeByte(s []byte) []byte {
	out := make([]byte, 0)
	pos1 := 0
	for {
		if pos1 >= len(s) {
			break
		}

		if isSymbol(s[pos1]) {
			out = append(out, s[pos1])
			pos1++
			continue
		}

		newPos1, newPos2, num := findBraces(s[pos1:])
		if newPos1 == -1 {
			out = append(out, s[pos1:]...)
			break
		}
		newPos1 += pos1
		newPos2 += pos1

		cp := makeCopy(decodeByte(s[newPos1+1:newPos2]), num)
		out = append(out, cp...)
		pos1 = newPos2 + 1
	}
	return out
}

func makeCopy(s []byte, i int) (out []byte) {
	if i <= 1 {
		return s
	}
	out = make([]byte, 0, len(s)*i)
	for range i {
		out = append(out, s...)
	}
	return out
}

func getNum(in byte) int {
	if in < 48 || in > 57 {
		return -1
	}
	return int(in - 48)
}

func isSymbol(in byte) bool {
	if in >= 97 && in <= 122 {
		return true
	}
	return false
}

func findBraces(s []byte) (pos1, pos2, num int) {
	count := 0
	for i := range s {
		if s[i] == '[' {
			if count == 0 {
				pos1 = i
			}
			count++
			continue
		}
		if count == 0 {
			n := getNum(s[i])
			if n >= 0 {
				if n == 0 {
				}
				num *= 10
				num += n
			}
			continue
		}
		if s[i] == ']' {
			count--
			if count == 0 {
				return pos1, i, num
			}
		}
	}
	return -1, -1, -1
}
