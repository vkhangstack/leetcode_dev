package main

import "fmt"

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	if s == goal {
		set := make(map[rune]struct{}, len(s))
		for _, v := range s {
			set[v] = struct{}{}
		}
		return len(set) < len(goal)
	}
	diff := make([]int, 0)
	for i := 0; i < len(s); i++ {
		if s[i] != goal[i] {
			diff = append(diff, i)
		}
	}

	return len(diff) == 2 && goal[diff[0]] == s[diff[1]] && s[diff[0]] == goal[diff[1]]
}

func main() {
	fmt.Println(buddyStrings("ab", "ba"))
}
