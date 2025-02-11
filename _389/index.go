package main

import "fmt"

func findTheDifference(s string, t string) byte {
	visit := make(map[rune]int, len(s))

	for _, v := range s {
		visit[v]++
	}
	for _, v := range t {
		if visit[v] == 0 {
			return byte(v)
		}
		visit[v]--
	}

	return byte(0)
}

func main() {
	fmt.Println(findTheDifference("abde", "abdef"))
}
