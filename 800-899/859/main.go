package main

func buddyStrings(s string, goal string) bool {
	if len(s) == 0 || len(s) == 1 || len(goal) == 0 || len(goal) == 1 {
		return false
	}
	if len(s) != len(goal) {
		return false
	}
	count := 0
	var s1, s2, g1, g2 byte
	for i := range s {
		if s[i] != goal[i] {
			if count == 0 {
				s1 = s[i]
				g1 = goal[i]
			} else {
				s2 = s[i]
				g2 = goal[i]
			}
			count++
			if count > 2 {
				return false
			}
		}
	}
	if count == 1 {
		return false
	}
	if count == 2 {
		if s1 == g2 && s2 == g1 {
			return true
		}
		return false
	}
	mp := make(map[byte]int)
	for i := range s {
		mp[s[i]]++
		val := mp[s[i]]
		if val == 2 {
			return true
		}
	}
	return false
}
