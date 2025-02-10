package firstuniqchar

func firstUniqChar(s string) int {
	//O(n+m)
	//visit := make(map[rune]int, len(s))
	//
	//for i := 0; i < len(s); i++ {
	//	visit[rune(s[i])] = visit[rune(s[i])] + 1
	//}
	//
	//for i, v := range s {
	//	if visit[v] == 1 {
	//		return i
	//	}
	//}
	//return -1

	set := [26]int{}
	for i := range len(s) {
		set[s[i]-'a']++
	}

	for i := range len(s) {
		if set[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}

func Run(s string) int {
	return firstUniqChar(s)
}
