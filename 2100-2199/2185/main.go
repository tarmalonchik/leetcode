package main

func prefixCount(words []string, pref string) int {
	var counter = 0
	for i := range words {
		if len(words[i]) < len(pref) {
			continue
		}
		for j := range pref {
			if pref[j] != words[i][j] {
				break
			}
			if j == len(pref)-1 {
				counter++
			}
		}
	}
	return counter
}
