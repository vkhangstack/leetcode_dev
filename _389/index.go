package main

import "fmt"

func findTheDifference(s string, t string) byte {
	//visit := make(map[rune]int, len(s))
	//
	//for _, v := range s {
	//	visit[v]++
	//}
	//for _, v := range t {
	//	if visit[v] == 0 {
	//		return byte(v)
	//	}
	//	visit[v]--
	//}
	//
	//return byte(0)

	// ways02 use two poiter

	sumS, sumT := 0, 0
	for _, v := range s {
		sumS += int(v)
	}
	for _, v := range t {
		sumT += int(v)
	}
	diff := sumT - sumS

	return byte(diff)
}

func main() {
	fmt.Println(findTheDifference("abcd", "abcde"))
}
