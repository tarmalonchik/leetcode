package main

func detectCapitalUse(word string) bool {
	if len(word) == 0 || len(word) == 1 {
		return true
	}

	allCap := false

	if isCap(word[1]) {
		if !isCap(word[0]) {
			return false
		}
		allCap = true
	}

	for i := 2; i < len(word); i++ {
		if isCap(word[i]) {
			if !allCap {
				return false
			}
		} else {
			if allCap {
				return false
			}
		}
	}
	return true
}

func isCap(a byte) bool {
	if a >= 65 && a <= 90 {
		return true
	}
	return false
}
