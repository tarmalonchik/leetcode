package main

import (
	"hash/maphash"
	"strconv"
)

type MyHashMap struct {
	seed maphash.Seed
	arr  []*list
}

func Constructor() MyHashMap {
	return MyHashMap{
		seed: maphash.MakeSeed(),
		arr:  make([]*list, 1000),
	}
}

func (this *MyHashMap) Put(key int, value int) {
	h := this.hash(key)
	if this.arr[h] == nil {
		this.arr[h] = &list{}
	}
	this.arr[h].add(key, value)
}

func (this *MyHashMap) Get(key int) int {
	h := this.hash(key)
	if this.arr[h] == nil {
		return -1
	}
	return this.arr[h].get(key)
}

func (this *MyHashMap) Remove(key int) {
	h := this.hash(key)
	if this.arr[h] == nil {
		return
	}
	this.arr[h].remove(key)
}

func (this *MyHashMap) hash(key int) int {
	num := maphash.Bytes(this.seed, []byte(strconv.Itoa(key)))
	return int(num % 1000)
}

type list struct {
	node *node
}

type node struct {
	key  int
	val  int
	next *node
}

func (l *list) get(key int) int {
	cur := l.node
	for {
		if cur == nil {
			return -1
		}
		if cur.key == key {
			return cur.val
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

func (l *list) add(key, val int) {
	if l.node == nil {
		l.node = &node{
			key: key,
			val: val,
		}
		return
	}

	cur := l.node
	for {
		if cur.key == key {
			cur.val = val
			return
		}
		if cur.next == nil {
			break
		}
		cur = cur.next
	}

	cur.next = &node{
		val: val,
		key: key,
	}
}
