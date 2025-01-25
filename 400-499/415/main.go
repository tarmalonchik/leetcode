package main

func addStrings(num1 string, num2 string) string {
	if len(num1) == 0 {
		return num2
	}
	if len(num2) == 0 {
		return num1
	}

	num1Custom := str(num1)
	num2Custom := str(num2)

	counter := 0
	prevSum := 0

	var outBytes = make([]byte, 0, maxNum(len(num1), len(num2)))

	for {
		int1, valid1 := num1Custom.getNum(counter)
		int2, valid2 := num2Custom.getNum(counter)

		if !valid1 && !valid2 {
			if prevSum != 0 {
				outBytes = append(outBytes, numToByte(prevSum))
			}
			break
		}

		localSum := int1 + int2 + prevSum
		if localSum >= 10 {
			localSum = localSum % 10
			prevSum = 1
		} else {
			prevSum = 0
		}
		outBytes = append(outBytes, numToByte(localSum))
		counter++
	}

	pos1 := 0
	pos2 := len(outBytes) - 1

	for {
		if pos2 <= pos1 {
			break
		}
		outBytes[pos1], outBytes[pos2] = outBytes[pos2], outBytes[pos1]
		pos1++
		pos2--
	}
	return string(outBytes)
}

type str string

func (s *str) getNum(idx int) (int, bool) {
	if idx > len(*s)-1 {
		return 0, false
	}
	return byteToNum((*s)[len(*s)-1-idx]), true
}

func byteToNum(in byte) int {
	return int(in - 48)
}

func numToByte(in int) byte {
	return byte(in + 48)
}

func maxNum(a, b int) int {
	if a < b {
		return b
	}
	return a
}
