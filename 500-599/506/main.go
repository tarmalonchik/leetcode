package main

import (
	"strconv"
)

func findRelativeRanks(score []int) []string {
	base := mergeSort(score)
	out := make([]string, len(score))

	counter := 0
	for i := len(base) - 1; i >= 0; i-- {
		out[base[i]] = getString(counter)
		counter++
	}
	return out
}

func getString(counter int) string {
	counter = counter + 1
	if counter == 1 {
		return "Gold Medal"
	} else if counter == 2 {
		return "Silver Medal"
	} else if counter == 3 {
		return "Bronze Medal"
	}
	return strconv.Itoa(counter)
}

func mergeSort(score []int) []int {
	base := make([]int, len(score))
	for i := range score {
		base[i] = i
	}
	_, base = recursion(score, base)
	return base
}

func recursion(score, base []int) ([]int, []int) {
	if len(score) == 0 || len(score) == 1 {
		return score, base
	}
	if len(score) == 2 {
		if score[0] > score[1] {
			score[1], score[0] = score[0], score[1]
			base[1], base[0] = base[0], base[1]
		}
		return score, base
	}

	out := make([]int, len(score))
	outBase := make([]int, len(score))
	one, oneBase := recursion(score[0:len(score)/2], base[0:len(base)/2])
	two, twoBase := recursion(score[len(score)/2:], base[len(base)/2:])

	pos1 := 0
	pos2 := 0
	pos := 0

	for {
		if pos1 == len(one) && pos2 == len(two) {
			break
		}
		if pos1 == len(one) || pos2 == len(two) {
			if pos1 < len(one) {
				out[pos] = one[pos1]
				outBase[pos] = oneBase[pos1]
				pos1++
				pos++
				continue
			} else {
				out[pos] = two[pos2]
				outBase[pos] = twoBase[pos2]
				pos2++
				pos++
				continue
			}
		}

		if one[pos1] < two[pos2] {
			out[pos] = one[pos1]
			outBase[pos] = oneBase[pos1]
			pos1++
			pos++
			continue
		} else {
			out[pos] = two[pos2]
			outBase[pos] = twoBase[pos2]
			pos2++
			pos++
			continue
		}
	}
	return out, outBase
}
