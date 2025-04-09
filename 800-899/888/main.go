package main

func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	aliceTotal := 0
	bobTotal := 0
	for i := range aliceSizes {
		aliceTotal += aliceSizes[i]
	}
	for i := range bobSizes {
		bobTotal += bobSizes[i]
	}

	diff := (aliceTotal - bobTotal) / 2
	mp := make(map[int]int)
	for i := range aliceSizes {
		mp[2*aliceSizes[i]-diff] = i
	}
	for i := range bobSizes {
		if val, ok := mp[2*bobSizes[i]+diff]; ok {
			return []int{aliceSizes[val], bobSizes[i]}
		}
	}
	return nil
}
