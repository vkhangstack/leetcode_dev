package main

import "fmt"

type MyHashSet struct {
	h map[int]struct{}
}

func Constructor() MyHashSet {
	return MyHashSet{
		h: make(map[int]struct{}),
	}
}

func (this *MyHashSet) Add(key int) {
	this.h[key] = struct{}{}
}

func (this *MyHashSet) Remove(key int) {
	if this.Contains(key) {
		delete(this.h, key)
	}
}

func (this *MyHashSet) Contains(key int) bool {
	if _, ok := this.h[key]; ok {
		return true
	}
	return false
}
func main() {
	key := 1
	obj := Constructor()
	obj.Add(key)
	obj.Remove(key)
	param_3 := obj.Contains(key)
	fmt.Println(param_3)
}
