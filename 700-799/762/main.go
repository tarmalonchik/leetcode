package main

func countPrimeSetBits(left int, right int) int {
	out := 0
	mp := map[int]interface{}{
		2:  nil,
		3:  nil,
		5:  nil,
		7:  nil,
		11: nil,
		13: nil,
		17: nil,
		19: nil,
	}
	some, ones := toArr(left)
	for i := left; i <= right; i++ {
		if _, ok := mp[ones]; ok {
			out++
		}
		some, ones = next(some, ones)
	}
	return out
}

func next(num []bool, ones int) ([]bool, int) {
	for i := range num {
		if num[i] {
			num[i] = false
			ones--
			if i == len(num)-1 {
				num = append(num, true)
				ones++
				break
			}
			continue
		}
		num[i] = true
		ones++
		break
	}
	return num, ones
}

func toArr(num int) ([]bool, int) {
	var out []bool
	var ones int
	for {
		if num == 0 {
			break
		}
		if num%2 == 0 {
			out = append(out, false)
		} else {
			ones++
			out = append(out, true)
		}
		num /= 2
	}
	return out, ones
}
