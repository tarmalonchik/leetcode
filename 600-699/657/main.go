package main

func judgeCircle(moves string) bool {
	vert := 0
	horiz := 0
	for i := range moves {
		switch moves[i] {
		case 'U':
			vert++
		case 'D':
			vert--
		case 'R':
			horiz++
		case 'L':
			horiz--
		}
	}
	return vert == 0 && horiz == 0
}
