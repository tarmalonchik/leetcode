package main

func multiply(num1 string, num2 string) string {
	if num1 == "" {
		return num2
	}
	if num2 == "" {
		return num1
	}
	out := ""

	if num1 == "0" || num2 == "0" {
		return "0"
	}

	pos := len(num1) - 1
	var resp []int
	for {
		if pos < 0 {
			break
		}
		localNum := multipleMultiDimSingleDim(num2, num1[pos])
		localNum = append(localNum, make([]int, len(num1)-1-pos)...)
		if resp == nil {
			resp = localNum
		} else {
			resp = sum(resp, localNum)
		}
		pos--
	}
	for i := range resp {
		out += numToString(resp[i])
	}
	return out
}

func sum(a, b []int) []int {
	out := make([]int, getMax(len(a), len(b))+1)
	posA := len(a) - 1
	posB := len(b) - 1
	outPos := 0
	mem := 0
	for {
		sumNum := 0
		if posA < 0 && posB < 0 {
			if mem > 0 {
				out[outPos] += mem
			}
			break
		} else if posA < 0 || posB < 0 {
			if posB < 0 {
				sumNum = a[posA]
			} else {
				sumNum = b[posB]
			}
		} else {
			sumNum = a[posA] + b[posB]
		}
		sumNum += mem
		out[outPos] += sumNum % 10
		mem = sumNum / 10
		outPos++
		posA--
		posB--
	}
	pos1 := 0
	pos2 := len(out) - 1
	for {
		if pos1 >= pos2 {
			break
		}
		out[pos1], out[pos2] = out[pos2], out[pos1]
		pos1++
		pos2--
	}
	if out[0] == 0 {
		return out[1:]
	}
	return out
}

func multipleMultiDimSingleDim(a string, b byte) []int {
	aPos := len(a) - 1
	mem := 0
	out := make([]int, len(a)+1)
	outPos := 0
	for {
		if aPos < 0 {
			if mem > 0 {
				out[outPos] += mem
			}
			break
		}
		mul := getNum(a[aPos]) * getNum(b)
		mul += mem
		out[outPos] += mul % 10
		mem = mul / 10
		outPos++
		aPos--
	}

	pos1 := 0
	pos2 := len(out) - 1
	for {
		if pos1 >= pos2 {
			break
		}
		out[pos1], out[pos2] = out[pos2], out[pos1]
		pos1++
		pos2--
	}
	if out[0] == 0 {
		return out[1:]
	}
	return out
}

func getNum(in byte) int {
	return int(in - 48)
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func numToString(in int) string {
	return string(byte(in + 48))
}
