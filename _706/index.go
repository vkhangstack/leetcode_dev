package main

import "fmt"

type MyHashMap struct {
	m map[int]int
}

func Constructor() MyHashMap {
	return MyHashMap{make(map[int]int)}
}

func (this *MyHashMap) Put(key int, value int) {
	this.m[key] = value
}

func (this *MyHashMap) Get(key int) int {
	if v, ok := this.m[key]; ok {
		return v
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	if _, ok := this.m[key]; ok {
		delete(this.m, key)
	}
}

func main() {
	obj := Constructor()
	obj.Put(1, 2)
	param2 := obj.Get(1)
	obj.Remove(1)
	fmt.Println(param2)

}
