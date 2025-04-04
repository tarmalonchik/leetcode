package main

func nextGreatestLetter(letters []byte, target byte) byte {
	pos1 := 0
	pos2 := len(letters) - 1

	for {
		if pos2-pos1 <= 1 {
			if letters[pos1] > target {
				return letters[pos1]
			}
			if letters[pos2] > target {
				return letters[pos2]
			}
			return letters[0]
		}
		center := pos1 + (pos2-pos1)/2
		if letters[center] <= target {
			pos1 = center
		} else {
			pos2 = center
		}
	}
}
