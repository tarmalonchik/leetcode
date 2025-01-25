package main

func longestPalindrome(s string) int {
	mp := make(map[byte]int)

	for i := range s {
		mp[s[i]]++
	}

	var out int
	var haveOne bool
	for _, val := range mp {
		if val == 1 {
			haveOne = true
			continue
		}
		if val >= 2 {
			out += (val / 2) * 2
			if val%2 == 1 {
				haveOne = true
			}
		}
	}
	if haveOne {
		out += 1
	}
	return out
}
