package main

func numJewelsInStones(jewels string, stones string) int {
	mp := make(map[byte]interface{})
	out := 0
	for i := range jewels {
		mp[jewels[i]] = nil
	}
	for i := range stones {
		if _, ok := mp[stones[i]]; ok {
			out++
		}
	}
	return out
}
