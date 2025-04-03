package main

func isOneBitCharacter(bits []int) bool {
	if len(bits) == 1 {
		return true
	}
	if bits[0] == 0 {
		return isOneBitCharacter(bits[1:])
	}
	if len(bits) == 2 {
		return false
	}
	return isOneBitCharacter(bits[2:])
}
