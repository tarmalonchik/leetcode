package main

func binaryGap(n int) int {
	if n < 2 {
		return 0
	}
	bin := toArr(n)

	pos1 := 0
	pos2 := 0
	for i := range bin {
		if bin[i] {
			pos1 = i
			pos2 = i + 1
			if pos2 > len(bin)-1 {
				return 0
			}
			break
		}
	}

	maxDist := 0
	for {
		if pos2 >= len(bin) {
			return maxDist
		}
		if bin[pos2] {
			if maxDist < pos2-pos1 {
				maxDist = pos2 - pos1
			}
			pos1 = pos2
			pos2 = pos1 + 1
			if pos2 > len(bin)-1 {
				return maxDist
			}
			continue
		}
		pos2++
	}
}

func toArr(num int) []bool {
	var out []bool
	for {
		if num == 0 {
			break
		}
		if num%2 == 0 {
			out = append(out, false)
		} else {
			out = append(out, true)
		}
		num /= 2
	}
	return out
}
