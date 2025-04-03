package main

import (
	"hash/maphash"
	"strconv"
)

type MyHashSet struct {
	seed maphash.Seed
	arr  []*list
}

func Constructor() MyHashSet {
	return MyHashSet{
		seed: maphash.MakeSeed(),
		arr:  make([]*list, 1000),
	}
}

func (this *MyHashSet) Add(key int) {
	h := this.hash(key)
	if this.arr[h] == nil {
		this.arr[h] = &list{}
	}
	this.arr[h].add(key)
}

func (this *MyHashSet) Remove(key int) {
	h := this.hash(key)
	if this.arr[h] == nil {
		return
	}
	this.arr[h].remove(key)
}

func (this *MyHashSet) Contains(key int) bool {
	h := this.hash(key)
	if this.arr[h] == nil {
		return false
	}
	return this.arr[h].contains(key)
}

func (this *MyHashSet) hash(key int) int {
	num := maphash.Bytes(this.seed, []byte(strconv.Itoa(key)))
	return int(num % 1000)
}

type list struct {
	node *node
}

type node struct {
	key  int
	next *node
}

func (l *list) contains(key int) bool {
	cur := l.node
	for {
		if cur == nil {
			return false
		}
		if cur.key == key {
			return true
		}
		cur = cur.next
	}
}

func (l *list) remove(key int) {
	if l.node == nil {
		return
	}

	var prev *node
	cur := l.node

	for {
		if cur.key == key {
			break
		}
		if cur.next == nil {
			return
		}
		prev = cur
		cur = cur.next
	}
	if cur.next == nil {
		if prev == nil {
			l.node = nil
			return
		}
		prev.next = nil
		return
	}
	if prev == nil {
		l.node = cur.next
		return
	}
	nx := cur.next
	prev.next = nx
}

func (l *list) add(key int) {
	if l.node == nil {
		l.node = &node{
			key: key,
		}
		return
	}

	cur := l.node
	for {
		if cur.key == key {
			return
		}
		if cur.next == nil {
			break
		}
		cur = cur.next
	}

	cur.next = &node{
		key: key,
	}
}
