package main

func countPrimes(n int) int {
	out := newPrimeNumbersCounter(n - 1)
	return len(out)
}

const noPreviousPrime = -2

// https://ru.wikipedia.org/wiki/Решето_Эратосфена
func newPrimeNumbersCounter(maxNum int) (out []int) {
	if maxNum <= 1 {
		return []int{}
	}

	out = make([]int, maxNum+1)
	out[0] = noPreviousPrime
	out[1] = noPreviousPrime

	currentIndex := 2
	for {
		out[currentIndex] = currentIndex

		for i := currentIndex + currentIndex; i < len(out); i += currentIndex {
			out[i] = noPreviousPrime
		}

		cont := false
		for i := currentIndex + 1; i < len(out); i++ {
			if out[i] == 0 {
				out[i] = i
				currentIndex = i
				cont = true
				break
			}
		}
		if cont {
			continue
		}
		break
	}

	currentIndex = 2
	for i := 2; i < len(out); i++ {
		if out[i] == noPreviousPrime {
			out[i] = currentIndex
		} else {
			currentIndex = i
		}
	}

	prevNum := noPreviousPrime
	var newOut = make([]int, 0)
	for i := range out {
		if out[i] == prevNum {
			continue
		}
		prevNum = out[i]
		newOut = append(newOut, prevNum)
	}

	return newOut
}
