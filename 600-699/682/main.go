package main

import (
	"strconv"
)

func calPoints(operations []string) int {
	p := processor{}
	for i := range operations {
		p.apply(operations[i])
	}
	return p.sum
}

type processor struct {
	values []int
	len    int
	sum    int
}

func (p *processor) apply(in string) {
	num, err := strconv.Atoi(in)
	if err == nil {
		p.add(num)
	} else {
		switch in {
		case "C":
			p.clearLast()
		case "D":
			p.double()
		case "+":
			p.plus()
		}
	}
}

func (p *processor) plus() {
	p.add(p.values[p.len-1] + p.values[p.len-2])
}

func (p *processor) clearLast() {
	p.sum -= p.values[p.len-1]
	p.len--
}

func (p *processor) double() {
	p.add(p.values[p.len-1] * 2)
}

func (p *processor) add(num int) {
	p.sum += num
	if len(p.values) <= p.len {
		p.values = append(p.values, num)
	} else {
		p.values[p.len] = num
	}
	p.len++
}
