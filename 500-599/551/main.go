package main

func checkRecord(s string) bool {
	sByte := []byte(s)
	absentCounter := 0
	lateCounter := 0

	for i := range sByte {
		if sByte[i] == 'A' {
			lateCounter = 0
			absentCounter++
			if absentCounter > 1 {
				return false
			}
		}
		if sByte[i] == 'L' {
			lateCounter++
			if lateCounter > 2 {
				return false
			}
		} else {
			lateCounter = 0
		}
	}
	return true
}
