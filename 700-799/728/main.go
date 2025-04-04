package main

func selfDividingNumbers(left int, right int) []int {
	out := make([]int, 0)
	nums, zero := numToArr(left - 1)
	for i := left; i <= right; i++ {
		nums, zero = next(nums, zero)
		if zero > 0 {
			continue
		}
		if divide(i, nums) {
			out = append(out, i)
		}
	}
	return out
}

func divide(num int, digs []int) bool {
	for i := range digs {
		if num%digs[i] == 0 {
			continue
		}
		return false
	}
	return true
}

func next(in []int, zeroCount uint8) ([]int, uint8) {
	for i := range in {
		if in[i] != 9 {
			if in[i] == 0 {
				zeroCount--
			}
			in[i]++
			break
		}
		in[i] = 0
		zeroCount++
		if i == len(in)-1 {
			in = append(in, 1)
			break
		}
	}
	return in, zeroCount
}

func numToArr(in int) ([]int, uint8) {
	if in == 0 {
		return []int{0}, 1
	}

	out := make([]int, 0)
	zeroCount := uint8(0)
	for {
		if in == 0 {
			break
		}
		num := in % 10
		if num == 0 {
			zeroCount++
		}
		out = append(out, num)
		in /= 10
	}
	return out, zeroCount
}
