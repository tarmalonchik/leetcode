package main

func flipAndInvertImage(image [][]int) [][]int {
	for i := range image {
		flipAndInvertRow(image[i])
	}
	return image
}

func flipAndInvertRow(in []int) {
	pos1 := 0
	pos2 := len(in) - 1

	for {
		if pos1 > pos2 {
			break
		}
		in[pos1], in[pos2] = in[pos2], in[pos1]
		if in[pos1] == 0 {
			in[pos1] = 1
		} else {
			in[pos1] = 0
		}
		if pos1 == pos2 {
			break
		}
		if in[pos2] == 0 {
			in[pos2] = 1
		} else {
			in[pos2] = 0
		}
		pos1++
		pos2--
	}
}
