package main

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	pos0 := 0
	pos1 := len(arr) - 1
	didChange := true
	for {
		if pos0 == pos1 {
			if pos0 != 0 && pos0 != len(arr)-1 {
				return true
			}
			return false
		}
		if !didChange {
			return false
		}
		didChange = false
		if arr[pos0+1] > arr[pos0] {
			didChange = true
			pos0++
		}
		if arr[pos1-1] > arr[pos1] {
			didChange = true
			pos1--
		}
	}
}
