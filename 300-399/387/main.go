package main

func firstUniqChar(s string) int {
	symbols := make(arr, 26)
	for i := len(s) - 1; i >= 0; i-- {
		symbols.setSymbol(s[i])
	}
	for i := range s {
		if symbols.getSymbol(s[i]) == 1 {
			return i
		}
	}
	return -1
}

type arr []int

func (a *arr) setSymbol(in byte) {
	in -= 97
	(*a)[in]++
}

func (a *arr) getSymbol(in byte) int {
	in -= 97
	return (*a)[in]
}
