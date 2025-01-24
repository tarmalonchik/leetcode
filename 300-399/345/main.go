package main

var mp = map[byte]interface{}{
	'a': nil,
	'e': nil,
	'i': nil,
	'o': nil,
	'u': nil,
	'A': nil,
	'E': nil,
	'I': nil,
	'O': nil,
	'U': nil,
}

func reverseVowels(s string) string {
	if len(s) == 0 {
		return s
	}

	sArr := make([]byte, len(s))
	for i := range s {
		sArr[i] = s[i]
	}

	pos1 := 0
	pos2 := len(s) - 1
	for {
		for i := pos1; i < len(s); i++ {
			if _, ok := mp[s[i]]; ok {
				pos1 = i
				break
			}
			if i == len(s)-1 {
				pos1 = len(s)
			}
		}
		for i := pos2; i >= 0; i-- {
			if _, ok := mp[s[i]]; ok {
				pos2 = i
				break
			}
			if i == 0 {
				pos2 = -1
			}
		}
		if pos1 >= pos2 {
			break
		}
		sArr[pos1], sArr[pos2] = sArr[pos2], sArr[pos1]
		pos1++
		pos2--
	}
	return string(sArr)
}
