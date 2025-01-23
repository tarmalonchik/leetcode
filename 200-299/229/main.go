package main

// https://en.wikipedia.org/wiki/Boyer–Moore_majority_vote_algorithm
// this algorithm for searching majority element with more than k/3 frequency
func majorityElement(nums []int) []int {
	set := 0
	a := 0
	b := 0
	aCount := 0
	bCount := 0

	for i := range nums {
		if set == 0 {
			a = nums[i]
			aCount++
			set++
			continue
		}
		if set == 1 {
			if nums[i] != a {
				b = nums[i]
				bCount++
				set++
				continue
			}
			aCount++
			continue
		}
		if nums[i] == a || nums[i] == b {
			if nums[i] == a {
				aCount++
				continue
			} else {
				bCount++
				continue
			}
		} else if aCount == 0 || bCount == 0 {
			if aCount == 0 {
				a = nums[i]
				aCount++
				continue
			} else {
				b = nums[i]
				bCount++
				continue
			}
		}
		bCount--
		aCount--
	}

	aRealCount := 0
	bRealCount := 0

	for i := range nums {
		if nums[i] == a && aCount > 0 {
			aRealCount++
		} else if nums[i] == b && bCount > 0 {
			bRealCount++
		}
	}

	out := make([]int, 0)
	if aRealCount > len(nums)/3 {
		out = append(out, a)
	}
	if bRealCount > len(nums)/3 {
		out = append(out, b)
	}
	return out
}
