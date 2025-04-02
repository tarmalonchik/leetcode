package main

func imageSmoother(img [][]int) [][]int {
	if len(img) == 0 || len(img[0]) == 0 {
		return img
	}

	upLine := make([]int, len(img[0]))
	copy(upLine, img[0])

	q := queue{}
	for i := range img[0] {
		orig := img[0][i]
		prevX, _ := q.checkLast()
		img[0][i] = calculate(nil, img[0], getLineIfPossible(img, 1), i, prevX)
		q.add(orig)
	}

	for y := 1; y < len(img); y++ {
		q = queue{}
		for x := 0; x < len(img[0]); x++ {
			orig := img[y][x]
			prevX, _ := q.checkLast()
			img[y][x] = calculate(upLine, img[y], getLineIfPossible(img, y+1), x, prevX)
			qOut, valid := q.add(orig)
			if valid {
				upLine[x-2] = qOut
			}
			if x == len(img[0])-1 {
				if x == 0 {
					upLine[0], _ = q.get()
				} else {
					upLine[x-1], _ = q.get()
					upLine[x], _ = q.get()
				}
			}
		}
	}
	return img
}

func getLineIfPossible(img [][]int, y int) []int {
	if y < 0 || y > len(img)-1 {
		return nil
	}
	return img[y]
}

func calculate(uLine, cLine, dLine []int, x int, leftXVal int) int {
	count := 1
	val := cLine[x]

	if x > 0 {
		count++
		val += leftXVal
	}
	if x < len(cLine)-1 {
		count++
		val += cLine[x+1]
	}

	if uLine != nil {
		count++
		val += uLine[x]
		if x > 0 {
			count++
			val += uLine[x-1]
		}
		if x < len(uLine)-1 {
			count++
			val += uLine[x+1]
		}
	}
	if dLine != nil {
		count++
		val += dLine[x]
		if x > 0 {
			count++
			val += dLine[x-1]
		}
		if x < len(dLine)-1 {
			count++
			val += dLine[x+1]
		}
	}
	return val / count
}

type queue struct {
	one      int
	oneValid bool
	two      int
	twoValid bool
}

func (s *queue) add(in int) (int, bool) {
	if !s.oneValid {
		s.one = in
		s.oneValid = true
		return 0, false
	}
	if !s.twoValid {
		s.two = in
		s.twoValid = true
		return 0, false
	}
	out := s.one
	s.one = s.two
	s.two = in
	return out, true
}

func (s *queue) get() (int, bool) {
	val, valid := s.one, s.oneValid
	s.one = s.two
	s.oneValid = s.twoValid
	s.twoValid = false
	return val, valid
}

func (s *queue) checkLast() (int, bool) {
	if s.twoValid {
		return s.two, s.twoValid
	}
	return s.one, s.oneValid
}
