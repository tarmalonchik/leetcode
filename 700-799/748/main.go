package main

func shortestCompletingWord(licensePlate string, words []string) string {
	origin, minSym, maxSym := createArr(licensePlate, 'a', 'z', true)
	minLen := -1
	minLenIdx := -1
	wordArr := make([]int, len(origin))

	for i := range words {
		if minLen >= 0 && len(words[i]) > minLen {
			continue
		}

		wordArr, _, _ = createArr(words[i], minSym, maxSym, false)
		compareResult := compare(origin, wordArr)
		for j := range wordArr {
			wordArr[j] = 0
		}
		if compareResult {
			if minLen < 0 {
				minLen = len(words[i])
				minLenIdx = i
				continue
			}
			if len(words[i]) < minLen {
				minLen = len(words[i])
				minLenIdx = i
			}
		}
	}
	return words[minLenIdx]
}

func createArr(word string, start, end byte, cutArr bool) (out []int, minSym, maxSym byte) {
	out = make([]int, end-start+1)
	minNum := uint8(0)
	maxNum := uint8(0)
	offset := start - 'a'
	minMaxSet := false
	for i := range word {
		if val, ok := charToNum(word[i]); ok {
			if numToChar(val) < start || numToChar(val) > end {
				continue
			}
			val -= offset
			if cutArr {
				if !minMaxSet {
					minNum = val
					maxNum = val
					minMaxSet = true
				} else {
					if val > maxNum {
						maxNum = val
					} else if val < minNum {
						minNum = val
					}
				}
			}
			out[val]++
		}
	}
	if cutArr {
		return out[minNum : maxNum+1], numToChar(minNum), numToChar(maxNum)
	}
	return out, 0, 0
}

func numToChar(num uint8) byte {
	return 97 + num
}

func compare(origin []int, cmp []int) bool {
	for i := range origin {
		if cmp[i] < origin[i] {
			return false
		}
	}
	return true
}

func charToNum(in byte) (uint8, bool) {
	if in >= 65 && in <= 90 {
		in += 32
	}
	if in >= 97 && in <= 122 {
		return in - 97, true
	}
	return 0, false
}
