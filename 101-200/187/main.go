package main

func findRepeatedDnaSequences(s string) []string {
	if len(s) < 10 {
		return nil
	}

	pos1 := 0
	pos2 := 9

	mp := make(map[string]int, len(s)-9)
	out := make([]string, 0)

	for {
		if pos2 >= len(s) {
			break
		}
		if val, ok := mp[s[pos1:pos2+1]]; ok {
			if val == 1 {
				out = append(out, s[pos1:pos2+1])
				mp[s[pos1:pos2+1]]++
			} else {
				mp[s[pos1:pos2+1]]++
			}
		} else {
			mp[s[pos1:pos2+1]] = 1
		}
		pos1++
		pos2++
	}
	return out
}
