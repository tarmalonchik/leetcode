package main

type RecentCounter struct {
	count int
	start *node
	end   *node
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	if this.start == nil || this.end == nil {
		this.start = &node{
			value: t,
		}
		this.end = this.start
		this.count = 1
		return 1
	}
	this.end.next = &node{
		value: t,
	}
	this.end = this.end.next
	this.count++

	t -= 3000
	for {
		if this.start.value >= t {
			break
		}
		this.count--
		this.start = this.start.next
	}
	return this.count
}

type node struct {
	value int
	next  *node
}
