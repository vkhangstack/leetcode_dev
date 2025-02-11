package _242

func isAnagram(s string, t string) bool {
	// ways 01
	if len(s) != len(t) {
		return false
	}

	visit := make(map[rune]int)

	for _, v := range s {
		visit[v]++
	}

	lettersToSatisfy := len(visit)

	for _, v := range t {
		if _, exist := visit[v]; !exist {
			return false
		}
		visit[v]--
		if visit[v] < 0 {
			return false
		}
		if visit[v] == 0 {
			lettersToSatisfy--
		}
	}
	return lettersToSatisfy == 0
}

func Run(s string, t string) bool {
	return isAnagram(s, t)
}
