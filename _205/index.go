package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	mapS := [256]byte{}
	mapsT := [256]byte{}

	for i := 0; i < len(s); i++ {
		if mapS[s[i]] != mapsT[t[i]] {
			return false
		}
		mapS[s[i]] = byte(i + 1)
		mapsT[t[i]] = byte(i + 1)
	}
	return true
	//mapS := map[uint8]int{}
	//mapT := map[uint8]int{}
	//for i := range s {
	//	fmt.Println(mapS)
	//	fmt.Println(mapT)
	//	if mapS[s[i]] != mapT[t[i]] {
	//		return false
	//	} else {
	//		mapS[s[i]] = i + 1
	//		mapT[t[i]] = i + 1
	//	}
	//}
	//return true

	//mapS := map[rune]int{}
	//mapsT := map[rune]int{}
	//for i := range s {
	//	if mapS[rune(s[i])] != mapsT[rune(t[i])] {
	//		return false
	//	} else {
	//		mapS[rune(s[i])] = i + 1
	//		mapsT[rune(t[i])] = i + 1
	//	}
	//
	//}
	//return true
}

func main() {
	fmt.Println(isIsomorphic("egg", "add"))
	fmt.Println(isIsomorphic("foo", "bar"))

}
