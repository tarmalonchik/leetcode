package main

func exist(board [][]byte, word string) bool {
	b := boarder{
		board:    board,
		occupied: make(mpManager),
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if b.recursion(word, 0, position{j, i}) {
				return true
			}
		}
	}
	return false
}

type boarder struct {
	board    [][]byte
	occupied mpManager
}

func (b *boarder) posValid(pos position) bool {
	if pos.x < 0 || pos.x > len(b.board[0])-1 {
		return false
	}
	if pos.y < 0 || pos.y > len(b.board)-1 {
		return false
	}
	return true
}

func (b *boarder) recursion(word string, wIdx int, pos position) bool {
	if !b.posValid(pos) {
		return false
	}

	if b.occupied.exists(pos) {
		return false
	}
	if wIdx == len(word)-1 {
		if b.board[pos.y][pos.x] == word[wIdx] {
			return true
		}
		return false
	}
	if b.board[pos.y][pos.x] != word[wIdx] {
		return false
	}
	b.occupied.add(pos)
	defer func() {
		b.occupied.del(pos)
	}()
	// up
	if b.recursion(word, wIdx+1, position{x: pos.x, y: pos.y - 1}) {
		return true
	}
	// down
	if b.recursion(word, wIdx+1, position{x: pos.x, y: pos.y + 1}) {
		return true
	}
	// right
	if b.recursion(word, wIdx+1, position{x: pos.x + 1, y: pos.y}) {
		return true
	}
	// left
	if b.recursion(word, wIdx+1, position{x: pos.x - 1, y: pos.y}) {
		return true
	}
	return false
}

type position struct {
	x, y int
}

type mpManager map[position]interface{}

func (m *mpManager) add(pos position) {
	(*m)[pos] = nil
}
func (m *mpManager) del(pos position) {
	delete(*m, pos)
}
func (m *mpManager) exists(pos position) bool {
	if _, ok := (*m)[pos]; ok {
		return true
	}
	return false
}
