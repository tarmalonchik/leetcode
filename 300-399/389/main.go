package main

func findTheDifference(s string, t string) byte {
	mp := make(map[byte]int)
	for i := range t {
		mp[t[i]]++
	}
	for i := range s {
		val, ok := mp[s[i]]
		if ok {
			if val == 1 {
				delete(mp, s[i])
			} else {
				mp[s[i]]--
			}
		}
	}
	for key, _ := range mp {
		return key
	}
	return 0
}
