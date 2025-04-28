package main

func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return nil
	}
	out := make([]int, 0)
	m := manager{
		state:  make(map[byte]int),
		origin: make(map[byte]int),
	}
	m.fillWithOrigin(p)
	for i := range s {
		head := i - len(p) + 1
		if m.add(s[i]) {
			out = append(out, head)
		}
		if head >= 0 {
			m.remove(s[head])
		}
	}
	return out
}

type manager struct {
	origin map[byte]int
	state  map[byte]int
	diff   int
}

func (m *manager) fillWithOrigin(origin string) {
	for i := range origin {
		m.origin[origin[i]]++
		m.diff++
	}
}

func (m *manager) add(symbol byte) bool {
	state := m.state[symbol]
	origin := m.origin[symbol]
	if mod(state, origin) > mod(state+1, origin) {
		m.diff--
	} else {
		m.diff++
	}
	m.state[symbol]++
	return m.diff == 0
}

func (m *manager) remove(symbol byte) bool {
	state := m.state[symbol]
	origin := m.origin[symbol]
	if mod(state, origin) > mod(state-1, origin) {
		m.diff--
	} else {
		m.diff++
	}
	m.state[symbol]--
	return m.diff == 0
}

func mod(a, b int) int {
	if a-b > 0 {
		return a - b
	}
	return b - a
}
