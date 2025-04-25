package main

type MyCalendar struct {
	list []interval
}

func Constructor() MyCalendar {
	return MyCalendar{
		list: make([]interval, 0, 1000),
	}
}

func (this *MyCalendar) Book(startTime int, endTime int) bool {
	in := interval{start: startTime, end: endTime}
	if len(this.list) == 0 {
		this.list = append(this.list, in)
		return true
	}
	pos, valid := this.find(in)
	if !valid {
		return false
	}

	this.list = append(this.list, interval{})
	next := this.list[pos]
	this.list[pos] = in

	for i := pos + 1; i < len(this.list); i++ {
		newNext := this.list[i]
		this.list[i] = next
		next = newNext
	}
	return true
}

func (this *MyCalendar) find(in interval) (int, bool) {
	pos1 := 0
	pos2 := len(this.list) - 1

	for {
		if pos2-pos1 <= 1 {
			if haveOverlap(this.list[pos1], in) || haveOverlap(this.list[pos2], in) {
				return 0, false
			}
			if in.start >= this.list[pos2].end {
				return pos2 + 1, true
			}
			if in.start >= this.list[pos1].end {
				return pos1 + 1, true
			}
			return pos1, true
		}
		center := pos1 + (pos2-pos1)/2

		if haveOverlap(this.list[center], in) {
			return 0, false
		}
		if in.start >= this.list[center].end {
			pos1 = center
		} else {
			pos2 = center
		}
	}
}

type interval struct {
	start int
	end   int
}

func haveOverlap(one, two interval) bool {
	if two.start >= one.end || one.start >= two.end {
		return false
	}
	return true
}
