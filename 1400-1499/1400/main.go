package main

func canConstruct(s string, k int) bool {
	mg := manager{
		mpSingle:   make(map[byte]int),
		mpHavePair: make(map[byte]int),
	}

	for i := range s {
		mg.addSymbol(s[i])
	}

	if len(mg.mpSingle) == k {
		return true
	}
	if len(mg.mpSingle) > k {
		return false
	}

	k = mg.decOdd(k)
	if k < 0 {
		return false
	}
	if k == 0 {
		return true
	}
	return mg.decRegular(k)
}

type manager struct {
	mpSingle     map[byte]int
	mpHavePair   map[byte]int
	oddCount     int
	regularCount int
}

func (m *manager) addSymbol(in byte) {
	if _, ok := m.mpSingle[in]; ok {
		delete(m.mpSingle, in)
		m.oddCount -= 1
		m.mpHavePair[in] += 2
		m.regularCount += 2
	} else {
		m.mpSingle[in] += 1
		m.oddCount += 1
	}
}

func (m *manager) decOdd(k int) int {
	if m.oddCount > k {
		return -1
	}
	k -= m.oddCount
	return k
}

func (m *manager) decRegular(k int) bool {
	if k%2 == 0 {
		if m.regularCount >= k {
			return true
		}
		return false
	}
	k--
	if m.regularCount > k {
		return true
	}
	return false
}
