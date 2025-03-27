package main

func licenseKeyFormatting(s string, k int) string {
	out := make([]byte, 0)

	newK := k
	i := len(s) - 1
	for {
		if i < 0 {
			break
		}
		if s[i] == '-' {
			i--
			continue
		}
		if newK == 0 {
			out = append(out, '-')
			newK = k
			continue
		}
		newK--
		out = append(out, lowerToUpper(s[i]))
		i--
	}

	pos1 := 0
	pos2 := len(out) - 1
	for {
		if pos1 >= pos2 {
			break
		}
		out[pos1], out[pos2] = out[pos2], out[pos1]
		pos1++
		pos2--
	}
	return string(out)
}

func lowerToUpper(in byte) byte {
	if in >= 97 && in <= 122 {
		return in - 32
	}
	return in
}
